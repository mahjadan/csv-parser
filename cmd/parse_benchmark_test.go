package cmd

import (
	"os"
	"rcsv/pkg/config"
	"rcsv/pkg/csvmapper"
	"rcsv/pkg/parser"
	"testing"
)

const inputCSVPath = "input/roster1.csv"

func BenchmarkParse(b *testing.B) {

	configLoader, err := config.NewConfigLoader(configPath)
	if err != nil {
		b.Fatalf("Error creating config loader? %v", err)
	}

	validFile, err := createTempCSVFile()
	if err != nil {
		b.Fatalf("Error creating temporary valid file: %v", err)
	}
	defer validFile.Close()

	invalidFile, err := createTempCSVFile()
	if err != nil {
		b.Fatalf("Error creating temporary invalid file: %v", err)
	}
	defer invalidFile.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		csvFile, err := os.Open(inputCSVPath)
		if err != nil {
			b.Fatalf("Error opening CSV file: %v", err)
		}
		defer csvFile.Close()

		columnIdentifier := csvmapper.NewDefaultColumnIdentifier()

		err = parser.Parse(csvFile, validFile, invalidFile, configLoader, columnIdentifier)
		if err != nil {
			b.Fatalf("Error running Parse: %v", err)
		}
	}

	b.StopTimer()
}

func createTempCSVFile() (*os.File, error) {
	file, err := os.CreateTemp("", "benchmark_*.csv")
	if err != nil {
		return nil, err
	}
	return file, nil
}
