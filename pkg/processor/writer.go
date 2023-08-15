package processor

// CSVWriter interface represents the methods used from csv.Writer
type CSVWriter interface {
	Write(record []string) error
}
