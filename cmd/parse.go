package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"rcsv/pkg/config"
	"rcsv/pkg/csvmapper"
	"rcsv/pkg/parser"
	"time"
)

var configPath = "config/config.json"
var outputDir = "output/"

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:     "parse [filename.csv]",
	Short:   "parse a given csv file.",
	Long:    `parse csv file and writes the valid and invalid records into separate files.`,
	Example: "rcsv parse input/roster1.csv",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runParse(args)
	},
}

func runParse(args []string) error {
	configLoader, err := config.NewConfigLoader(configPath)
	if err != nil {
		return err
	}
	csvFileName := args[0]
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
	defer printFileSize(validFile)
	defer printFileSize(invalidFile)
	defer printFileSize(csvFile)
	return parser.Parse(csvFile, validFile, invalidFile, configLoader, columnIdentifier)
}

func printFileSize(validFile *os.File) {
	stat, err := validFile.Stat()
	if err == nil {
		fmt.Printf("%v file size: %v bytes \n", stat.Name(), stat.Size())
	}
}

func createCSVFile(fileNamePrefix string) (*os.File, error) {
	fileName := generateFileName(fileNamePrefix)
	fullPath := filepath.Join(outputDir, fileName)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return nil, errors.Wrapf(err, "error creating directory %s", outputDir)
	}
	file, err := os.Create(fullPath)
	if err != nil {
		return nil, errors.Wrapf(err, "error creating %s CSV file", fileNamePrefix)
	}
	return file, nil
}
func generateFileName(prefix string) string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s_%s.csv", prefix, timestamp)
}

func init() {
	parseCmd.Flags().StringVarP(&configPath, "config", "c", configPath, "Path to the config file (optional)")
	parseCmd.Flags().StringVarP(&outputDir, "output", "o", outputDir, "Path to the output directory (optional)")
	rootCmd.AddCommand(parseCmd)
}
