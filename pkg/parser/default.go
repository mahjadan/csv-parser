package parser

import (
	"employee-csv-parser/pkg/config"
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

func Parse(csvFile io.Reader, columnsConfig *config.Loader, columnIdentifier csvmapper.ColumnIdentifier) error {
	reader := csv.NewReader(csvFile)
	headers, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV headers:", err)
		return errors.Wrap(err, "error reading CSV headers")
	}
	headers = utils.ToLowerTrimSlice(headers)

	err = columnIdentifier.MapColumnToIndexes(headers, columnsConfig.ColumnAliasConfig)
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
	defer validWriter.Flush()
	defer invalidWriter.Flush()
	csvProcessor := processor.NewCSVProcessor(validWriter, invalidWriter, columnIdentifier)

	err = csvProcessor.WriteHeaders(columnsConfig.ValidColumnNames, columnsConfig.InvalidColumnNames)
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
