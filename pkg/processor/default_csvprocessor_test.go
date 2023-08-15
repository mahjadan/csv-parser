package processor_test

import (
	"employee-csv-parser/pkg/csvmapper"
	"employee-csv-parser/pkg/processor"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultCSVProcessor(t *testing.T) {

	t.Run("ProcessValidRecord", func(t *testing.T) {
		mockWriter, mockInvalidWriter := mockWriters(t)
		columnIdentifier := mockColumnIdentifier(t)

		processor := processor.NewCSVProcessor(mockWriter, mockInvalidWriter, &columnIdentifier)
		record := []string{"1", "John", "john@example.com", "5000"}

		processor.ProcessValidRecord(record)

		mockWriter.AssertExpectations(t)
		mockInvalidWriter.AssertNotCalled(t, "Write", mock.Anything)
		assert.Equal(t, 1, mockWriter.WriteCallCount)
		assert.Equal(t, record, mockWriter.LastWriteArgument)
	})

	t.Run("ProcessValidRecord_InvalidEmployee", func(t *testing.T) {
		mockWriter, mockInvalidWriter := mockWriters(t)
		columnIdentifier := mockColumnIdentifier(t)

		processor := processor.NewCSVProcessor(mockWriter, mockInvalidWriter, &columnIdentifier)
		record := []string{"1", "John", "john@", "$90"}

		processor.ProcessValidRecord(record)

		mockInvalidWriter.AssertExpectations(t)
		mockWriter.AssertNotCalled(t, "Write", mock.Anything)
		assert.Equal(t, 1, mockInvalidWriter.WriteCallCount)
		assert.Equal(t, append(record, "Invalid email format"), mockInvalidWriter.LastWriteArgument)
	})

	t.Run("WriteHeaders", func(t *testing.T) {
		mockWriter, mockInvalidWriter := mockWriters(t)
		columnIdentifier := mockColumnIdentifier(t)

		processor := processor.NewCSVProcessor(mockWriter, mockInvalidWriter, &columnIdentifier)
		validHeaders := []string{"id", "name", "email"}
		invalidHeaders := []string{"id", "name", "email", "errors"}

		err := processor.WriteHeaders(validHeaders, invalidHeaders)

		mockWriter.AssertExpectations(t)
		mockInvalidWriter.AssertExpectations(t)
		assert.NoError(t, err)
		assert.Equal(t, 1, mockWriter.WriteCallCount)
		assert.Equal(t, 1, mockInvalidWriter.WriteCallCount)
		assert.Equal(t, validHeaders, mockWriter.WriteCalls[0])
		assert.Equal(t, invalidHeaders, mockInvalidWriter.WriteCalls[0])
	})
}

func mockColumnIdentifier(t *testing.T) csvmapper.MockColumnIdentifier {
	t.Helper()
	columnIdentifier := csvmapper.MockColumnIdentifier{}
	columnIdentifier.On("IndexForColumn", "id").Return(0)
	columnIdentifier.On("IndexForColumn", "name").Return(1)
	columnIdentifier.On("IndexForColumn", "email").Return(2)
	columnIdentifier.On("IndexForColumn", "salary").Return(3)
	return columnIdentifier
}

func mockWriters(t *testing.T) (*processor.MockCSVWriter, *processor.MockCSVWriter) {
	t.Helper()
	mockWriter := &processor.MockCSVWriter{}
	mockInvalidWriter := &processor.MockCSVWriter{}
	mockWriter.On("Write", mock.Anything).Return(nil).Once()
	mockInvalidWriter.On("Write", mock.Anything).Return(nil).Once()
	return mockWriter, mockInvalidWriter
}
