package encryption

import (
	"bytes"
	"errors"
	"testing"

	"github.com/dzibukalexander/file-processing/internal/config"
	"github.com/dzibukalexander/file-processing/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockEncryptor is a mock for the Encryptor interface
type MockEncryptor struct {
	mock.Mock
}

func (m *MockEncryptor) Encrypt(data []byte, key []byte) ([]byte, error) {
	args := m.Called(data, key)
	return args.Get(0).([]byte), args.Error(1)
}

// MockDecryptor is a mock for the Decryptor interface
type MockDecryptor struct {
	mock.Mock
}

func (m *MockDecryptor) Decrypt(data []byte, key []byte) ([]byte, error) {
	args := m.Called(data, key)
	return args.Get(0).([]byte), args.Error(1)
}

func setupLoggingTest() *bytes.Buffer {
	config.AppConfig = &config.Config{EnableLogging: true}
	logger.SetupLogger()

	logOutput := new(bytes.Buffer)
	logger.GetInstance().SetOutput(logOutput)
	return logOutput
}

func TestLoggingEncryptor(t *testing.T) {
	logOutput := setupLoggingTest()
	originalOutput := logger.GetInstance().Out
	defer logger.GetInstance().SetOutput(originalOutput)

	mockEnc := new(MockEncryptor)

	data := []byte("data")
	key := []byte("key")
	mockEnc.On("Encrypt", data, key).Return([]byte("encrypted"), nil)

	loggingEnc := NewLoggingEncryptor(mockEnc)
	_, err := loggingEnc.Encrypt(data, key)

	assert.NoError(t, err)
	assert.Contains(t, logOutput.String(), "Starting encryption")
	assert.Contains(t, logOutput.String(), "Encryption finished")
	mockEnc.AssertExpectations(t)
}

func TestLoggingDecryptor_Error(t *testing.T) {
	logOutput := setupLoggingTest()
	originalOutput := logger.GetInstance().Out
	defer logger.GetInstance().SetOutput(originalOutput)

	mockDec := new(MockDecryptor)

	data := []byte("data")
	key := []byte("key")
	expectedErr := errors.New("decryption error")
	mockDec.On("Decrypt", data, key).Return([]byte(nil), expectedErr)

	loggingDec := NewLoggingDecryptor(mockDec)
	_, err := loggingDec.Decrypt(data, key)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Contains(t, logOutput.String(), "Starting decryption")
	assert.Contains(t, logOutput.String(), "Decryption failed")
	mockDec.AssertExpectations(t)
}
