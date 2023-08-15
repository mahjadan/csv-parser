package csvmapper

type ColumnIdentifier interface {
	IndexForColumn(columnName string) int
	MapColumnToIndexes(csvHeaders []string, columnAliases map[string][]string) error
}
