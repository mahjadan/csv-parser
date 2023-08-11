/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"employee-csv-parser/pkg/parser"
	"employee-csv-parser/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

// var cfg config.ColumnNames
var cfg map[string][]string

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
		configFile, err := os.Open(configPath)
		if err != nil {
			return errors.Wrapf(err, "opening config file")
		}
		defer configFile.Close()

		decoder := json.NewDecoder(configFile)
		err = decoder.Decode(&cfg)
		if err != nil {

			return errors.Wrapf(err, "error decoding config file [%v]", configPath)
		}
		// todo remove log
		fmt.Printf("%+v\n", cfg)
		//normalizeConfig(&cfg)
		normalizeConfigMap(cfg)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		csvFile := args[0]
		fmt.Println("Config:", cfg)

		err := parser.Parse(cfg, csvFile)

		return err
	},
}

func normalizeConfigMap(cfg map[string][]string) {
	for k, v := range cfg {
		cfg[k] = utils.ToLowerTrimSlice(v)
	}
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
