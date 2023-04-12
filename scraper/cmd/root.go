package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	showSlug  string
	showShort string
	showImg   string

	skipMetadata bool
)

var rootCmd = &cobra.Command{
	Use:   "scraper",
	Short: "Scrape different services for show episodes",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if !skipMetadata {
			if showSlug == "" || showImg == "" {
				return fmt.Errorf("--skip-metadata so --show-slug and --show-img must be set as well")
			}
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(peacockCmd)
	rootCmd.AddCommand(netflixCmd)

	rootCmd.PersistentFlags().BoolVar(&skipMetadata, "skip-metadata", false, "Skip generating and writing the metadata.json file. This will ")
	rootCmd.PersistentFlags().StringVar(&showShort, "show-short", "", "Show short name. Like 'theofficeus'")
	rootCmd.PersistentFlags().StringVar(&showSlug, "show-slug", "", "Show slug. Like 'The Office US'")
	rootCmd.PersistentFlags().StringVar(&showImg, "show-img", "", "A link to the shows poster. Like 'https://www.amazon.com/bribase-shop-Michael-Office-poster/dp/B07G76KVGS'")
	rootCmd.MarkPersistentFlagRequired("show-short")
}
