/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"rcsv/pkg/config"
	"rcsv/pkg/csvmapper"
	"rcsv/pkg/parser"
	"time"
)

var columnsConfig *config.Loader

const configPath = "config/config.json"

var StandardOutputColumns = []string{"id", "name", "email", "salary"}
var InvalidOutputColumns = append(StandardOutputColumns, "errors")

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
			return errors.Wrapf(err, "opening config file [%v]", configPath)
		}
		defer configFile.Close()

		ld := config.NewConfigLoader(configFile, StandardOutputColumns, InvalidOutputColumns)
		err = ld.LoadConfig()
		if err != nil {
			return err
		}
		columnsConfig = ld
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		csvFileName := args[0]
		fmt.Println("Run Config:", columnsConfig)
		csvFile, err := os.Open(csvFileName)
		if err != nil {
			return errors.Wrap(err, "error opening CSV file")
		}
		defer csvFile.Close()
		validFile, err := createCSVFile("valid")
		if err != nil {
			return err
		}
		defer validFile.Close()

		invalidFile, err := createCSVFile("invalid")
		if err != nil {
			return err
		}
		defer invalidFile.Close()

		columnIdentifier := csvmapper.NewDefaultColumnIdentifier()
		err = parser.Parse(csvFile, validFile, invalidFile, columnsConfig, columnIdentifier)
		stat, err := validFile.Stat()
		if err == nil {
			fmt.Printf("Valid file size: %v bytes \n", stat.Size())
		}
		istat, err := invalidFile.Stat()
		if err == nil {
			fmt.Printf("Invalid file size: %v bytes \n", istat.Size())
		}

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
func createCSVFile(fileNamePrefix string) (*os.File, error) {
	fileName := generateFileName(fileNamePrefix)
	file, err := os.Create(fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "error creating %s CSV file", fileNamePrefix)
	}
	return file, nil
}

func generateFileName(prefix string) string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s_%s.csv", prefix, timestamp)
}
