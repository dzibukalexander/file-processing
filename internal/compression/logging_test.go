package compression

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

// MockCompressor is a mock for the Compressor interface
type MockCompressor struct {
	mock.Mock
}

func (m *MockCompressor) Compress(data []byte) ([]byte, error) {
	args := m.Called(data)
	return args.Get(0).([]byte), args.Error(1)
}

// MockDecompressor is a mock for the Decompressor interface
type MockDecompressor struct {
	mock.Mock
}

func (m *MockDecompressor) Decompress(data []byte) ([]byte, error) {
	args := m.Called(data)
	return args.Get(0).([]byte), args.Error(1)
}

func setupLoggingTest() *bytes.Buffer {
	config.AppConfig = &config.Config{EnableLogging: true}
	logger.SetupLogger()

	logOutput := new(bytes.Buffer)
	logger.GetInstance().SetOutput(logOutput)
	return logOutput
}

func TestLoggingCompressor(t *testing.T) {
	runner.Run(t, "LoggingCompressor", func(t provider.T) {
		t.WithNewStep("success path", func(s provider.StepCtx) {
			logOutput := setupLoggingTest()
			origOut := logger.GetInstance().Out
			defer logger.GetInstance().SetOutput(origOut)

			mockComp := new(MockCompressor)
			input := []byte("data")
			mockComp.On("Compress", input).Return([]byte("compressed"), nil)

			loggingComp := NewLoggingCompressor(mockComp)
			_, err := loggingComp.Compress(input)

			s.Assert().NoError(err)
			s.Assert().Contains(logOutput.String(), "Starting compression")
			s.Assert().Contains(logOutput.String(), "Compression finished")
			mockComp.AssertExpectations(t)
		})
	})
}

func TestLoggingDecompressor_Error(t *testing.T) {
	runner.Run(t, "LoggingDecompressor", func(t provider.T) {
		t.WithNewStep("error path", func(s provider.StepCtx) {
			logOutput := setupLoggingTest()
			origOut := logger.GetInstance().Out
			defer logger.GetInstance().SetOutput(origOut)

			mockDecomp := new(MockDecompressor)
			input := []byte("data")
			expectedErr := errors.New("decompression error")
			mockDecomp.On("Decompress", input).Return([]byte(nil), expectedErr)

			loggingDecomp := NewLoggingDecompressor(mockDecomp)
			_, err := loggingDecomp.Decompress(input)

			s.Assert().Error(err)
			s.Assert().Equal(expectedErr, err)
			s.Assert().Contains(logOutput.String(), "Starting decompression")
			s.Assert().Contains(logOutput.String(), "Decompression failed")
			mockDecomp.AssertExpectations(t)
		})
	})
}
