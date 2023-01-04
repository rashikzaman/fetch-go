package cmd

import (
	"fmt"
	"os"
	"url-saver/app/scrapper"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch is a web page downloader",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			url := args[0]
			err := scrapper.GetHtmlFromUrl(url)
			if err != nil {
				fmt.Println("error running command", err)
			}
		} else {
			fmt.Println("Please provide an url")
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
