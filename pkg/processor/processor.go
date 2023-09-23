package processor

type CSVProcessor interface {
	Flush()
	InitializeHeaders(validColumnNames []string, headers []string) error
	ProcessRecord(record []string, err error) bool
	GetUniqueRecords() int64
}
