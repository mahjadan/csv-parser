package processor_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	mock2 "rcsv/pkg/mock"
	"rcsv/pkg/processor"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultCSVProcessor(t *testing.T) {

	t.Run("ProcessValidRecord", func(t *testing.T) {
		mockWriter, mockInvalidWriter := mockWriters(t)
		columnIdentifier := mockColumnIdentifier(t)

		processor := processor.NewCSVProcessor(mockWriter, mockInvalidWriter, columnIdentifier)
		record := []string{"1", "John", "john@example.com", "5000"}

		validRecord := processor.ProcessRecord(record, nil)

		mockWriter.AssertExpectations(t)
		mockInvalidWriter.AssertNotCalled(t, "Write", mock.Anything)
		assert.Equal(t, 1, mockWriter.WriteCallCount)
		assert.Equal(t, record, mockWriter.LastWriteArgument)
		assert.True(t, validRecord)
	})
	t.Run("ProcessValidRecord_WithDuplicatedEmail", func(t *testing.T) {
		mockWriter, mockInvalidWriter := mockWriters(t)
		columnIdentifier := mockColumnIdentifier(t)

		processor := processor.NewCSVProcessor(mockWriter, mockInvalidWriter, columnIdentifier)
		record := []string{"1", "John", "john@example.com", "5000"}

		processor.ProcessRecord(record, nil)
		processor.ProcessRecord(record, nil)

		mockWriter.AssertNumberOfCalls(t, "Write", 1)
		assert.Equal(t, 1, mockWriter.WriteCallCount)
		assert.Equal(t, record, mockWriter.LastWriteArgument)
	})

	t.Run("ProcessValidRecord_WithDifferentColumnOrder", func(t *testing.T) {
		mockWriter, mockInvalidWriter := mockWriters(t)
		columnIdentifier := &mock2.MockColumnIdentifier{}
		columnIdentifier.On("IndexForColumn", "id").Return(2)
		columnIdentifier.On("IndexForColumn", "name").Return(1)
		columnIdentifier.On("IndexForColumn", "email").Return(3)
		columnIdentifier.On("IndexForColumn", "salary").Return(0)
		processor := processor.NewCSVProcessor(mockWriter, mockInvalidWriter, columnIdentifier)
		record := []string{"5000", "John", "1", "john@example.com"}

		processor.ProcessRecord(record, nil)

		expected := []string{"1", "John", "john@example.com", "5000"}
		mockWriter.AssertExpectations(t)
		mockInvalidWriter.AssertNotCalled(t, "Write", mock.Anything)
		assert.Equal(t, 1, mockWriter.WriteCallCount)
		assert.Equal(t, expected, mockWriter.LastWriteArgument)
	})

	t.Run("ProcessValidRecord_InvalidEmployee", func(t *testing.T) {
		mockWriter, mockInvalidWriter := mockWriters(t)
		columnIdentifier := mockColumnIdentifier(t)

		processor := processor.NewCSVProcessor(mockWriter, mockInvalidWriter, columnIdentifier)
		record := []string{"1", "John", "john@", "$90"}

		validRecord := processor.ProcessRecord(record, nil)

		mockInvalidWriter.AssertExpectations(t)
		mockWriter.AssertNotCalled(t, "Write", mock.Anything)
		assert.Equal(t, 1, mockInvalidWriter.WriteCallCount)
		assert.Equal(t, append(record, "Invalid email format"), mockInvalidWriter.LastWriteArgument)
		assert.False(t, validRecord)
	})

	t.Run("ProcessValidRecord_InvalidEmployee", func(t *testing.T) {
		mockWriter, mockInvalidWriter := mockWriters(t)
		columnIdentifier := mockColumnIdentifier(t)

		processor := processor.NewCSVProcessor(mockWriter, mockInvalidWriter, columnIdentifier)
		record := []string{"1", "John", "john@"}

		validRecord := processor.ProcessRecord(record, errors.New("missing colum data"))

		mockInvalidWriter.AssertExpectations(t)
		mockWriter.AssertNotCalled(t, "Write", mock.Anything)
		assert.Equal(t, 1, mockInvalidWriter.WriteCallCount)
		assert.Equal(t, append(record, "missing colum data"), mockInvalidWriter.LastWriteArgument)
		assert.False(t, validRecord)
	})

	t.Run("InitializeHeaders", func(t *testing.T) {
		mockWriter, mockInvalidWriter := mockWriters(t)
		columnIdentifier := mockColumnIdentifier(t)

		processor := processor.NewCSVProcessor(mockWriter, mockInvalidWriter, columnIdentifier)
		validHeaders := []string{"id", "name", "email"}
		invalidHeaders := []string{"id", "name", "email", "errors"}

		err := processor.InitializeHeaders(validHeaders, invalidHeaders)

		mockWriter.AssertExpectations(t)
		mockInvalidWriter.AssertExpectations(t)
		assert.NoError(t, err)
		assert.Equal(t, 1, mockWriter.WriteCallCount)
		assert.Equal(t, 1, mockInvalidWriter.WriteCallCount)
		assert.Equal(t, validHeaders, mockWriter.WriteCalls[0])
		assert.Equal(t, invalidHeaders, mockInvalidWriter.WriteCalls[0])
	})
	t.Run("Flush", func(t *testing.T) {
		mockWriter := &mock2.CSVWriter{}
		mockInvalidWriter := &mock2.CSVWriter{}
		mockWriter.On("Flush").Return(nil)
		mockInvalidWriter.On("Flush").Return(nil)

		columnIdentifier := mockColumnIdentifier(t)
		processor := processor.NewCSVProcessor(mockWriter, mockInvalidWriter, columnIdentifier)

		processor.Flush()

		mockWriter.AssertExpectations(t)
		assert.Equal(t, 1, mockWriter.FlushCallCount)
		assert.Equal(t, 1, mockInvalidWriter.FlushCallCount)
	})
}

func mockColumnIdentifier(t *testing.T) *mock2.MockColumnIdentifier {
	t.Helper()
	columnIdentifier := mock2.MockColumnIdentifier{}
	columnIdentifier.On("IndexForColumn", "id").Return(0)
	columnIdentifier.On("IndexForColumn", "name").Return(1)
	columnIdentifier.On("IndexForColumn", "email").Return(2)
	columnIdentifier.On("IndexForColumn", "salary").Return(3)
	return &columnIdentifier
}

func mockWriters(t *testing.T) (*mock2.CSVWriter, *mock2.CSVWriter) {
	t.Helper()
	mockWriter := &mock2.CSVWriter{}
	mockInvalidWriter := &mock2.CSVWriter{}
	mockWriter.On("Write", mock.Anything).Return(nil)
	mockInvalidWriter.On("Write", mock.Anything).Return(nil)
	return mockWriter, mockInvalidWriter
}
