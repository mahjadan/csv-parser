package parser

import (
	"encoding/csv"
	"io"
	"rcsv/pkg/config"
	"rcsv/pkg/csvmapper"
	"rcsv/pkg/processor"
	"rcsv/pkg/utils"

	"github.com/pkg/errors"
)

type ParseStats struct {
	TotalInvalidRecords int64
	TotalValidRecords   int64
	UniqueRecords       int64
}

func Parse(csvFile io.Reader, configLoader config.Mapper, columnIdentifier csvmapper.ColumnIdentifier, csvProcessor processor.CSVProcessor) (*ParseStats, error) {
	var stats ParseStats

	reader := csv.NewReader(csvFile)
	headers, err := reader.Read()
	if err != nil {
		return nil, errors.Wrap(err, "error reading CSV headers")
	}
	headers = utils.ToLowerTrimSlice(headers)

	err = columnIdentifier.MapColumnToIndexes(headers, configLoader.GetColumnAliasMap())
	if err != nil {
		return nil, err
	}

	defer csvProcessor.Flush()
	err = csvProcessor.InitializeHeaders(configLoader.GetValidColumnNames(), append(headers, "errors"))
	if err != nil {
		return nil, errors.Wrap(err, "error writing headers")
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if valid := csvProcessor.ProcessRecord(record, err); valid {
			stats.TotalValidRecords++
		} else {
			stats.TotalInvalidRecords++
		}
	}

	stats.UniqueRecords = csvProcessor.GetUniqueRecords()
	return &stats, nil
}
