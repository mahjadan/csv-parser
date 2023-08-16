package parser

import (
	"employee-csv-parser/pkg/config"
	"employee-csv-parser/pkg/csvmapper"
	"employee-csv-parser/pkg/processor"
	"employee-csv-parser/pkg/utils"
	"encoding/csv"
	"github.com/pkg/errors"
	"io"
)

func Parse(csvFile io.Reader, validFile, invalidFile io.Writer, columnsConfig *config.Loader, columnIdentifier csvmapper.ColumnIdentifier) error {
	reader := csv.NewReader(csvFile)
	headers, err := reader.Read()
	if err != nil {
		return errors.Wrap(err, "error reading CSV headers")
	}
	headers = utils.ToLowerTrimSlice(headers)

	err = columnIdentifier.MapColumnToIndexes(headers, columnsConfig.ColumnAliasConfig)
	if err != nil {
		return err
	}

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
