package parser

import (
	"employee-csv-parser/pkg/csvmapper"
	"employee-csv-parser/pkg/processor"
	"employee-csv-parser/pkg/utils"
	"encoding/csv"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"time"
)

func Parse(columnAliases map[string][]string, csfFilePath string) error {
	csvFile, err := os.Open(csfFilePath)
	if err != nil {
		return errors.Wrap(err, "error opening CSV file")
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	headers, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV headers:", err)
		return errors.Wrap(err, "error reading CSV headers")
	}
	headers = utils.ToLowerTrimSlice(headers)

	columnIdentifier := csvmapper.NewColumnIdentifier(columnAliases, headers)
	err = columnIdentifier.MapColumnToIndexes(headers)
	if err != nil {
		return err
	}

	// create writers
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

	validWriter := csv.NewWriter(validFile)
	invalidWriter := csv.NewWriter(invalidFile)
	// move flush inside the processor
	defer validWriter.Flush()
	defer invalidWriter.Flush()
	// todo why the name CSVProcessor ? its specific to csv ?
	csvProcessor := processor.NewCSVProcessor(validWriter, invalidWriter, columnIdentifier)
	// Write headers to the valid CSV file
	err = csvProcessor.InitializeHeaders()
	if err != nil {
		return errors.Wrap(err, "error writing headers")
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			csvProcessor.WriteInvalidRecord(record, err.Error())
			continue
		}
		csvProcessor.ProcessValidRecord(record)
	}
	return nil
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
