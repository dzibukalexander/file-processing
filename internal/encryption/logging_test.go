package encryption

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
	runner.Run(t, "LoggingEncryptor", func(t provider.T) {
		t.WithNewStep("success path", func(s provider.StepCtx) {
			logOutput := setupLoggingTest()
			origOut := logger.GetInstance().Out
			defer logger.GetInstance().SetOutput(origOut)

			mockEnc := new(MockEncryptor)
			data := []byte("data")
			key := []byte("key")
			mockEnc.On("Encrypt", data, key).Return([]byte("encrypted"), nil)

			loggingEnc := NewLoggingEncryptor(mockEnc)
			_, err := loggingEnc.Encrypt(data, key)

			s.Assert().NoError(err)
			s.Assert().Contains(logOutput.String(), "Starting encryption")
			s.Assert().Contains(logOutput.String(), "Encryption finished")
			mockEnc.AssertExpectations(t)
		})
	})
}

func TestLoggingDecryptor_Error(t *testing.T) {
	runner.Run(t, "LoggingDecryptor", func(t provider.T) {
		t.WithNewStep("error path", func(s provider.StepCtx) {
			logOutput := setupLoggingTest()
			origOut := logger.GetInstance().Out
			defer logger.GetInstance().SetOutput(origOut)

			mockDec := new(MockDecryptor)
			data := []byte("data")
			key := []byte("key")
			expectedErr := errors.New("decryption error")
			mockDec.On("Decrypt", data, key).Return([]byte(nil), expectedErr)

			loggingDec := NewLoggingDecryptor(mockDec)
			_, err := loggingDec.Decrypt(data, key)

			s.Assert().Error(err)
			s.Assert().Equal(expectedErr, err)
			s.Assert().Contains(logOutput.String(), "Starting decryption")
			s.Assert().Contains(logOutput.String(), "Decryption failed")
			mockDec.AssertExpectations(t)
		})
	})
}
