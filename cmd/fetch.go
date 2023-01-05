package cmd

import (
	"fetch/app/scrapper"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var includeMeta bool

var rootCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch is a web page downloader",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, arg := range args {
				url := arg
				metaData, err := scrapper.GetHtmlFromUrl(url, includeMeta)
				if err != nil {
					fmt.Println("error running command", err)
				}
				if metaData != nil {
					fmt.Println("site", metaData.Site)
					fmt.Println("num_links", metaData.NumLinks)
					fmt.Println("images", metaData.Images)
					fmt.Println("last_fetch", metaData.LastFetch.Format(time.RFC850))
				}
			}
		} else {
			fmt.Println("Please provide atleast one url")
			return
		}
	},
}

func Execute() {
	rootCmd.Flags().BoolVarP(&includeMeta, "metadata", "m", false, "include metadata for webpage basic information")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
