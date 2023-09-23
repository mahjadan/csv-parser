package csvmapper

import (
	"github.com/pkg/errors"
	"slices"
	"strings"
)

type DefaultColumnIdentifier struct {
	columnIndexMap map[string]int
}

func (c *DefaultColumnIdentifier) IndexForColumn(columnName string) int {
	return c.columnIndexMap[columnName]
}
func NewDefaultColumnIdentifier() *DefaultColumnIdentifier {
	return &DefaultColumnIdentifier{
		columnIndexMap: make(map[string]int),
	}
}
func (c *DefaultColumnIdentifier) MapColumnToIndexes(csvHeaders []string, columnAliases map[string][]string) error {
	for columnName := range columnAliases {
		c.columnIndexMap[columnName] = -1
	}

	for index, header := range csvHeaders {
		for columnName, alternativeNames := range columnAliases {
			if header == strings.ToLower(strings.TrimSpace(columnName)) || slices.Contains(alternativeNames, header) {
				c.columnIndexMap[columnName] = index
			}
		}
	}
	return c.hasMissingColumn(columnAliases)
}

func (c *DefaultColumnIdentifier) hasMissingColumn(columnAliases map[string][]string) error {
	for columnName, index := range c.columnIndexMap {
		if index == -1 {
			return errors.Errorf("missing column/header: '%s'. Tried alternatives: %v", columnName, columnAliases[columnName])
		}
	}
	return nil
}
