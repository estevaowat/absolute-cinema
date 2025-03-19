package cmd

import (
	"fmt"
	"os"

	"github.com/estevaowat/absolute-cinema/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "absci",
	Short: "generate and save movies in database",
	Run: func(cmd *cobra.Command, args []string) {
		println("printing using root cmd")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	api.GetMovies(100)
}
