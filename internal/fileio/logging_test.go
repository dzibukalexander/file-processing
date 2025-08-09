package fileio

import (
	"bytes"
	"errors"
	"testing"

	"github.com/dzibukalexander/file-processing/internal/config"
	"github.com/dzibukalexander/file-processing/internal/logger"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
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
	runner.Run(t, "LoggingFileReader", func(t provider.T) {
		t.WithNewStep("success path", func(s provider.StepCtx) {
			logOutput := setupLoggingTest()
			origOut := logger.GetInstance().Out
			defer logger.GetInstance().SetOutput(origOut)

			mockReader := new(MockFileReader)
			path := "file.txt"
			mockReader.On("Read", path).Return([]byte("content"), nil)

			loggingReader := NewLoggingFileReader(mockReader)
			_, err := loggingReader.Read(path)

			s.Assert().NoError(err)
			s.Assert().Contains(logOutput.String(), "Starting file read")
			s.Assert().Contains(logOutput.String(), "File read finished")
			mockReader.AssertExpectations(t)
		})
	})
}

func TestLoggingFileWriter_Error(t *testing.T) {
	runner.Run(t, "LoggingFileWriter", func(t provider.T) {
		t.WithNewStep("error path", func(s provider.StepCtx) {
			logOutput := setupLoggingTest()
			origOut := logger.GetInstance().Out
			defer logger.GetInstance().SetOutput(origOut)

			mockWriter := new(MockFileWriter)
			path := "file.txt"
			data := []byte("content")
			expectedErr := errors.New("write error")
			mockWriter.On("Write", path, data).Return(expectedErr)

			loggingWriter := NewLoggingFileWriter(mockWriter)
			err := loggingWriter.Write(path, data)

			s.Assert().Error(err)
			s.Assert().Equal(expectedErr, err)
			s.Assert().Contains(logOutput.String(), "Starting file write")
			s.Assert().Contains(logOutput.String(), "File write failed")
			mockWriter.AssertExpectations(t)
		})
	})
}
