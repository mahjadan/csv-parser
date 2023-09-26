package cmd

import (
	"encoding/csv"
	"os"
	"rcsv/pkg/config"
	"rcsv/pkg/csvmapper"
	"rcsv/pkg/parser"
	"rcsv/pkg/processor"
	"testing"
)

const inputCSVPath = "input/data1.csv"

func BenchmarkParse(b *testing.B) {

	configLoader, err := config.NewConfigLoader(configPath)
	if err != nil {
		b.Fatalf("Error creating config loader? %v", err)
	}

	validFile, err := createTempCSVFile()
	if err != nil {
		b.Fatalf("Error creating temporary valid file: %v", err)
	}

	invalidFile, err := createTempCSVFile()
	if err != nil {
		b.Fatalf("Error creating temporary invalid file: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		csvFile, err := os.Open(inputCSVPath)
		if err != nil {
			b.Fatalf("Error opening CSV file: %v", err)
		}
		defer csvFile.Close()

		columnIdentifier := csvmapper.NewDefaultColumnIdentifier()

		csvProcessor := processor.NewCSVProcessor(validFile, invalidFile, columnIdentifier)
		_, err = parser.Parse(csvFile, configLoader, columnIdentifier, csvProcessor)
		if err != nil {
			b.Fatalf("Error running Parse: %v", err)
		}
	}
	b.StopTimer()
}

func createTempCSVFile() (*csv.Writer, error) {
	file, err := os.CreateTemp("", "benchmark_*.csv")
	csvFile := csv.NewWriter(file)
	if err != nil {
		return nil, err
	}
	return csvFile, nil
}
