package mock

import testifymock "github.com/stretchr/testify/mock"

type CSVWriter struct {
	testifymock.Mock
	WriteCallCount    int
	WriteCalls        [][]string
	LastWriteArgument []string
	FlushCallCount    int
}

func (m *CSVWriter) Flush() {
	m.Called()
	m.FlushCallCount++
}

func (m *CSVWriter) Write(record []string) error {
	args := m.Called(record)
	m.WriteCallCount++
	m.WriteCalls = append(m.WriteCalls, record)
	m.LastWriteArgument = record
	return args.Error(0)
}
