/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/estevaowat/absolute-cinema/api"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("getting flag type", cmd.Flag("type").Value.String())
		log.Println("args", args)

		if cmd.Flag("type").Value.String() != "csv" {
			log.Fatal("only csv type implemented!")
		}

		length, err := strconv.Atoi(cmd.Flag("length").Value.String())

		if err != nil {
			log.Fatal("length invalid, should be more than zero")
		}

		if length <= 0 {
			log.Fatal("length is invalid, should be more than zero")
		}

		goroutine, err := strconv.ParseBool(cmd.Flag("goroutine").Value.String())
		if err != nil {
			log.Fatal("could not parse goroutine", goroutine)
		}

		if goroutine {
			api.GetMoviesUsingGoRoutines(length)
		} else {
			api.GetMovies(length)
		}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	rootCmd.PersistentFlags().StringP("type", "t", "csv", "default file is csv")
	rootCmd.PersistentFlags().IntP("length", "l", 10, "number of lines of file")
	rootCmd.PersistentFlags().Bool("goroutine", false, "run command using goroutines")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
