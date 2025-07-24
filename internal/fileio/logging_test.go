package fileio

import (
	"bytes"
	"errors"
	"testing"

	"github.com/dzibukalexander/file-processing/internal/config"
	"github.com/dzibukalexander/file-processing/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockFileReader is a mock for the FileReader interface
type MockFileReader struct {
	mock.Mock
}

func (m *MockFileReader) Read(filePath string) ([]byte, error) {
	args := m.Called(filePath)
	return args.Get(0).([]byte), args.Error(1)
}

// MockFileWriter is a mock for the FileWriter interface
type MockFileWriter struct {
	mock.Mock
}

func (m *MockFileWriter) Write(filePath string, data []byte) error {
	args := m.Called(filePath, data)
	return args.Error(0)
}

func setupLoggingTest() *bytes.Buffer {
	config.AppConfig = &config.Config{EnableLogging: true}
	logger.SetupLogger()

	logOutput := new(bytes.Buffer)
	logger.GetInstance().SetOutput(logOutput)
	return logOutput
}

func TestLoggingFileReader(t *testing.T) {
	logOutput := setupLoggingTest()
	originalOutput := logger.GetInstance().Out
	defer logger.GetInstance().SetOutput(originalOutput)

	mockReader := new(MockFileReader)

	path := "file.txt"
	mockReader.On("Read", path).Return([]byte("content"), nil)

	loggingReader := NewLoggingFileReader(mockReader)
	_, err := loggingReader.Read(path)

	assert.NoError(t, err)
	assert.Contains(t, logOutput.String(), "Starting file read")
	assert.Contains(t, logOutput.String(), "File read finished")
	mockReader.AssertExpectations(t)
}

func TestLoggingFileWriter_Error(t *testing.T) {
	logOutput := setupLoggingTest()
	originalOutput := logger.GetInstance().Out
	defer logger.GetInstance().SetOutput(originalOutput)

	mockWriter := new(MockFileWriter)

	path := "file.txt"
	data := []byte("content")
	expectedErr := errors.New("write error")
	mockWriter.On("Write", path, data).Return(expectedErr)

	loggingWriter := NewLoggingFileWriter(mockWriter)
	err := loggingWriter.Write(path, data)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Contains(t, logOutput.String(), "Starting file write")
	assert.Contains(t, logOutput.String(), "File write failed")
	mockWriter.AssertExpectations(t)
}
