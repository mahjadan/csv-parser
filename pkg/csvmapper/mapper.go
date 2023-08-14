package csvmapper

import (
	"employee-csv-parser/pkg/utils"
	"github.com/pkg/errors"
)

type ColumnIdentifier struct {
	columnIndexMap     map[string]int
	columnAliases      map[string][]string
	ValidColumnNames   []string
	InvalidColumnNames []string
}

var standardOutputColumns = []string{"id", "name", "email", "salary"}

func (c *ColumnIdentifier) IndexForColumn(columnName string) int {
	return c.columnIndexMap[columnName]
}
func NewColumnIdentifier(columnAliases map[string][]string, csvHeaders []string) ColumnIdentifier {
	columnIndexMap := make(map[string]int)
	for columnName := range columnAliases {
		columnIndexMap[columnName] = -1
	}
	return ColumnIdentifier{
		columnIndexMap:     columnIndexMap,
		columnAliases:      columnAliases,
		InvalidColumnNames: append(csvHeaders, "errors"),
		ValidColumnNames:   standardOutputColumns,
	}
}
func (c *ColumnIdentifier) MapColumnToIndexes(csvHeaders []string) error {
	for index, header := range csvHeaders {
		for columnName, alternativeNames := range c.columnAliases {
			if utils.SliceContains(alternativeNames, header) {
				c.columnIndexMap[columnName] = index
			}
		}
	}
	return c.hasMissingColumn()
}

func (c *ColumnIdentifier) hasMissingColumn() error {
	for columnName, index := range c.columnIndexMap {
		if index == -1 {
			return errors.Errorf("missing column/header: '%s'. Tried alternatives: %v", columnName, c.columnAliases[columnName])
		}
	}
	return nil
}
