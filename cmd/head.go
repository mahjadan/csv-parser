/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// headCmd represents the head command
var headCmd = &cobra.Command{
	Use:   "head",
	Short: "show the first 3 lines of the csv file",
	Long: `This helps in identifying the headers of the csv file that need to be parsed
and usage of using your command. For example:

c-parser.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("head called")
	},
}

func init() {
	rootCmd.AddCommand(headCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// headCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// headCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
