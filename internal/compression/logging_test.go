package compression

import (
	"bytes"
	"errors"
	"testing"

	"github.com/dzibukalexander/file-processing/internal/config"
	"github.com/dzibukalexander/file-processing/internal/logger"
	"github.com/stretchr/testify/assert"
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
	// Resetting the output is handled by each test
	return logOutput
}

func TestLoggingCompressor(t *testing.T) {
	logOutput := setupLoggingTest()
	originalOutput := logger.GetInstance().Out
	defer logger.GetInstance().SetOutput(originalOutput)

	mockComp := new(MockCompressor)

	input := []byte("data")
	mockComp.On("Compress", input).Return([]byte("compressed"), nil)

	loggingComp := NewLoggingCompressor(mockComp)
	_, err := loggingComp.Compress(input)

	assert.NoError(t, err)
	assert.Contains(t, logOutput.String(), "Starting compression")
	assert.Contains(t, logOutput.String(), "Compression finished")
	mockComp.AssertExpectations(t)
}

func TestLoggingDecompressor_Error(t *testing.T) {
	logOutput := setupLoggingTest()
	originalOutput := logger.GetInstance().Out
	defer logger.GetInstance().SetOutput(originalOutput)

	mockDecomp := new(MockDecompressor)

	input := []byte("data")
	expectedErr := errors.New("decompression error")
	mockDecomp.On("Decompress", input).Return([]byte(nil), expectedErr)

	loggingDecomp := NewLoggingDecompressor(mockDecomp)
	_, err := loggingDecomp.Decompress(input)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Contains(t, logOutput.String(), "Starting decompression")
	assert.Contains(t, logOutput.String(), "Decompression failed")
	mockDecomp.AssertExpectations(t)
}
