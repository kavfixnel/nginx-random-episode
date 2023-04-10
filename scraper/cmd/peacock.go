package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/integralist/go-findroot/find"
	"github.com/spf13/cobra"
)

var (
	showPath  string
	showSlug  string
	showShort string
	showImg   string

	skipMetadata bool
)

// peacockCmd represents the base command when called without any subcommands
var peacockCmd = &cobra.Command{
	Use:   "peacock",
	Short: "Scrape Peacock for a show",
	PreRunE: func(cmd *cobra.Command, args []string) error {

		if !skipMetadata {
			if showSlug == "" || showImg == "" {
				return fmt.Errorf("--skip-metadata so --show-slug and --show-img must be set as well")
			}
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		root, err := find.Repo()
		if err != nil {
			return err
		}

		req, err := http.NewRequest("GET", fmt.Sprintf("https://atom.peacocktv.com/adapter-calypso/v3/query/node?slug=%s&represent=(items(items))", showPath), nil)
		if err != nil {
			return err
		}
		req.Header.Add("x-skyott-proposition", "NBCUOTT")
		req.Header.Add("x-skyott-territory", "US")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var peacockResp peacockResp
		json.Unmarshal([]byte(body), &peacockResp)

		outputFiles := make(map[int]*os.File)
		defer func() {
			for _, of := range outputFiles {
				if err := of.Close(); err != nil {
					panic(err)
				}
			}
		}()

		os.MkdirAll(fmt.Sprintf("%s/episodes/peacock/%s", root.Path, showShort), os.ModePerm)
		allEpisodesFile, err := os.Create(fmt.Sprintf("%s/episodes/peacock/%s/all.episodes", root.Path, showShort))
		if err != nil {
			return err
		}
		defer allEpisodesFile.Close()

		for _, show := range peacockResp.Relationships.Items.Data {
			for _, episode := range show.Relationships.Items.Data {
				if _, ok := outputFiles[episode.Attributes.SeasonNumber]; !ok {
					of, err := os.Create(fmt.Sprintf("%s/episodes/peacock/%s/s%d.episodes", root.Path, showShort, episode.Attributes.SeasonNumber))
					if err != nil {
						return err
					}

					outputFiles[episode.Attributes.SeasonNumber] = of
				}

				episodeLink := fmt.Sprintf("https://www.peacocktv.com/watch/playback/vod/%s/%s\n",
					episode.Attributes.Formats.HD.ContentID,
					episode.Attributes.ProviderVariantID,
				)

				outputFiles[episode.Attributes.SeasonNumber].WriteString(episodeLink)
				allEpisodesFile.WriteString(episodeLink)
			}
		}

		if !skipMetadata {
			response, err := http.Get(showImg)
			if err != nil {
				return err
			}
			defer response.Body.Close()

			file, err := os.Create(fmt.Sprintf("%s/static/images/%s.jpeg", root.Path, showShort))
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(file, response.Body)
			if err != nil {
				return err
			}

			// Write metadata.json file
			metadataFile, err := os.Create(fmt.Sprintf("%s/episodes/peacock/%s/metadata.json", root.Path, showShort))
			if err != nil {
				return err
			}
			defer metadataFile.Close()

			metadataContent := showMetadata{
				Slug:        showSlug,
				Imgref:      fmt.Sprintf("/images/%s.jpeg", showShort),
				ShowPath:    showPath,
				OverviewUrl: "https://www.peacocktv.com/watch/asset" + showPath,
				NumSeasons:  len(outputFiles),
			}
			metadata, err := json.MarshalIndent(metadataContent, "", "    ")
			if err != nil {
				return err
			}
			metadataFile.Write(metadata)
		}

		return nil
	},
}

func init() {
	peacockCmd.Flags().StringVar(&showPath, "show-path", "", "Show path in Peacock. Like '/tv/parks-and-recreation/5883799404534408112'")
	peacockCmd.Flags().StringVar(&showShort, "show-short", "", "Show short name. Like 'theofficeus'")
	peacockCmd.MarkFlagRequired("show-path")
	peacockCmd.MarkFlagRequired("show-short")

	peacockCmd.Flags().BoolVar(&skipMetadata, "skip-metadata", false, "Skip generating and writing the metadata.json file. This will ")
	peacockCmd.Flags().StringVar(&showSlug, "show-slug", "", "Show slug. Like 'The Office US'")
	peacockCmd.Flags().StringVar(&showImg, "show-img", "", "A link to the shows poster. Like 'https://www.amazon.com/bribase-shop-Michael-Office-poster/dp/B07G76KVGS'")

	// peacockCmd.MarkFlagRequired("show-slug")
	// peacockCmd.MarkFlagRequired("show-img")
}
