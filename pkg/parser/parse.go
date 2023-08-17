package parser

import (
	"encoding/csv"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"rcsv/pkg/config"
	"rcsv/pkg/csvmapper"
	"rcsv/pkg/processor"
	"rcsv/pkg/utils"
	"time"
)

func Parse(csvFile io.Reader, validFile, invalidFile io.Writer, configLoader *config.Loader, columnIdentifier csvmapper.ColumnIdentifier) error {
	startTime := time.Now()
	var totalInvalidRecords int64 = 0
	var totalValidRecords int64 = 0

	reader := csv.NewReader(csvFile)
	headers, err := reader.Read()
	if err != nil {
		return errors.Wrap(err, "error reading CSV headers")
	}
	headers = utils.ToLowerTrimSlice(headers)

	err = columnIdentifier.MapColumnToIndexes(headers, configLoader.ColumnAliasConfig)
	if err != nil {
		return err
	}

	validWriter := csv.NewWriter(validFile)
	invalidWriter := csv.NewWriter(invalidFile)
	defer validWriter.Flush()
	defer invalidWriter.Flush()
	csvProcessor := processor.NewCSVProcessor(validWriter, invalidWriter, columnIdentifier)

	err = csvProcessor.WriteHeaders(configLoader.ValidColumnNames, append(headers, "errors"))
	if err != nil {
		return errors.Wrap(err, "error writing headers")
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			csvProcessor.WriteInvalidRecord(record, err.Error())
			totalInvalidRecords++
			continue
		}
		csvProcessor.ProcessValidRecord(record)
		totalValidRecords++
	}
	endTime := time.Now()
	processingTime := endTime.Sub(startTime)

	fmt.Println("No of Invalid records: ", totalInvalidRecords)
	fmt.Println("No of Valid records: ", totalValidRecords)
	fmt.Println("No of Unique records: ", csvProcessor.GetUniqueRecords())
	fmt.Println("Processing Time: ", processingTime)
	return nil
}
