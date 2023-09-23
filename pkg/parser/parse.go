package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"rcsv/pkg/config"
	"rcsv/pkg/csvmapper"
	"rcsv/pkg/processor"
	"rcsv/pkg/utils"
	"time"

	"github.com/pkg/errors"
)

func Parse(csvFile io.Reader, configLoader config.Mapper, columnIdentifier csvmapper.ColumnIdentifier, csvProcessor processor.CSVProcessor) error {
	startTime := time.Now()
	var totalInvalidRecords int64 = 0
	var totalValidRecords int64 = 0

	//todo: check if this need to be refactored why creating csv reader here?
	reader := csv.NewReader(csvFile)
	headers, err := reader.Read()
	if err != nil {
		return errors.Wrap(err, "error reading CSV headers")
	}
	headers = utils.ToLowerTrimSlice(headers)

	err = columnIdentifier.MapColumnToIndexes(headers, configLoader.GetColumnAliasMap())
	if err != nil {
		return err
	}

	defer csvProcessor.Flush()
	err = csvProcessor.InitializeHeaders(configLoader.GetValidColumnNames(), append(headers, "errors"))
	if err != nil {
		return errors.Wrap(err, "error writing headers")
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if valid := csvProcessor.ProcessRecord(record, err); valid {
			totalValidRecords++
		} else {
			totalInvalidRecords++
		}
	}

	endTime := time.Now()
	processingTime := endTime.Sub(startTime)

	//todo: find better way instead of printing here
	fmt.Println("No of Invalid records: ", totalInvalidRecords)
	fmt.Println("No of Valid records: ", totalValidRecords)
	fmt.Println("No of Unique records: ", csvProcessor.GetUniqueRecords())
	fmt.Println("Processing Time: ", processingTime)
	return nil
}
