package processor

type CSVProcessor interface {
	ProcessValidRecord(record []string)
	WriteHeaders(valid, invalid []string) error
}
