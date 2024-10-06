package cmd

import (
	"ofa/cmd/ad"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ofa",
	Short: "one command line tool for all the things",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(ad.NewADCmd())
}
