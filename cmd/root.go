package cmd

import (
	"log"
	"strconv"

	"github.com/estevaowat/absolute-cinema/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "absci",
	Short: "generate and save movies in database",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("printing using root cmd using args=", args)
		length, err := strconv.Atoi(args[0])

		if err != nil {
			log.Println("could not parse args", err)
		}

		api.GetMoviesUsingGoRoutines(length)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	}
}
