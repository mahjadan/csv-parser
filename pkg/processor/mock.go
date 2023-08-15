package processor

import "github.com/stretchr/testify/mock"

type MockCSVWriter struct {
	mock.Mock
	WriteCallCount    int
	WriteCalls        [][]string
	LastWriteArgument []string
}

func (m *MockCSVWriter) Write(record []string) error {
	args := m.Called(record)
	m.WriteCallCount++
	m.WriteCalls = append(m.WriteCalls, record)
	m.LastWriteArgument = record
	return args.Error(0)
}
