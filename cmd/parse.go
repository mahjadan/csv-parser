/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"employee-csv-parser/pkg/config"
	"employee-csv-parser/pkg/parser"
	"fmt"
	"github.com/spf13/cobra"
)

var columnConfig map[string][]string

const configPath = "config/config.json"

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse [filename.csv]",
	Short: "parse a given csv file.",
	Long: `parse employees csv file and writes the valid and invalid records into separate files
and usage of using your command. For example:
rcsv parse source/roster1.csv.`,
	Example: "rcsv parse input/roster1.csv",
	Args:    cobra.ExactArgs(1),

	PreRunE: func(cmd *cobra.Command, args []string) error {
		loader := config.NewConfigLoader(configPath)

		cfg, err := loader.LoadConfig()
		if err != nil {
			return err
		}

		columnConfig = cfg
		// todo remove log
		fmt.Printf("PreRunE config %+v\n", columnConfig)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		csvFile := args[0]
		fmt.Println("Run Config:", columnConfig)
		err := parser.Parse(columnConfig, csvFile)
		return err
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
