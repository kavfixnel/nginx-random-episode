package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	log "github.com/inconshreveable/log15"
	"github.com/integralist/go-findroot/find"
	"github.com/spf13/cobra"
)

var (
	showId     string
	numSeasons int
)

func getSeasonUrlFromImdb(show string, season int) ([]string, error) {
	url := fmt.Sprintf("https://www.imdb.com/title/%s/episodes/_ajax/?season=%d", show, season)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return []string{""}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return []string{""}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []string{""}, err
	}

	re := regexp.MustCompile(`/title/tt[0-9]+`)

	episodeDictionary := make(map[string]struct{}, 0)
	for _, v := range re.FindAll(body, -1) {
		episodeDictionary[string(v)] = struct{}{}
	}

	episodeList := make([]string, 0, len(episodeDictionary))
	for k := range episodeDictionary {
		episodeList = append(episodeList, k[7:])
	}

	return episodeList, nil
}

func getNetflixUrlFromImdb(showId string) (string, error) {
	url := "https://www.imdb.com/watch/_ajax/option"
	method := "POST"

	payload := strings.NewReader("minibar=" + showId)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var imdbShowResp imdbShowResp
	json.Unmarshal(body, &imdbShowResp)

	re := regexp.MustCompile(`https://www.netflix.com/watch/\d+`)

	for _, v := range imdbShowResp.Minibar {
		return string(re.Find([]byte(v))), nil
	}

	return "", fmt.Errorf("could not find link in request for %s", showId)
}

// netflixCmd represents the base command when called without any subcommands
var netflixCmd = &cobra.Command{
	Use:   "netflix",
	Short: "Scrape Netflix for a show",
	RunE: func(cmd *cobra.Command, args []string) error {
		srvlog := log.New("service", "netflix", "show", showShort)

		root, err := find.Repo()
		if err != nil {
			return err
		}

		os.MkdirAll(fmt.Sprintf("%s/episodes/netflix/%s", root.Path, showShort), os.ModePerm)
		allEpisodesFile, err := os.Create(fmt.Sprintf("%s/episodes/netflix/%s/all.episodes", root.Path, showShort))
		if err != nil {
			return err
		}
		defer allEpisodesFile.Close()

		totalEpisodes := 0
		for i := 1; i <= numSeasons; i++ {
			episodes, err := getSeasonUrlFromImdb(showId, i)
			if err != nil {
				return err
			}

			of, err := os.Create(fmt.Sprintf("%s/episodes/netflix/%s/s%d.episodes", root.Path, showShort, i))
			if err != nil {
				return err
			}

			seasonEpisodes := 0
			for _, episode := range episodes {
				link, err := getNetflixUrlFromImdb(episode)
				if err != nil {
					continue
				}

				if link != "" {
					of.WriteString(link + "\n")
					allEpisodesFile.WriteString(link + "\n")
					totalEpisodes++
					seasonEpisodes++
				}
			}

			srvlog.Info("Finished processing season", "season", i, "episodes", seasonEpisodes)

			of.Close()
		}

		if !skipMetadata {
			srvlog.Debug("Writing metadata file")
			err := generateMetadata(
				root.Path,
				"netflix",
				fmt.Sprintf("/images/%s.jpeg", showShort),
				"",
				numSeasons,
			)
			if err != nil {
				return err
			}
		}

		srvlog.Info("Finished", "seasons", numSeasons, "episodes", totalEpisodes)

		return nil
	},
}

func init() {
	netflixCmd.Flags().StringVar(&showId, "showId", "", "Show ID on IMDb (from URL). Like tt0903747 in 'https://www.imdb.com/title/tt0903747/'")
	netflixCmd.Flags().IntVar(&numSeasons, "num-seasons", 0, "Number of seasons starting from 1")
	netflixCmd.MarkFlagRequired("showId")
	netflixCmd.MarkFlagRequired("num-seasons")
}
