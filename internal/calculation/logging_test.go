package calculation

import (
	"bytes"
	"errors"
	"testing"

	"github.com/dzibukalexander/file-processing/internal/config"
	"github.com/dzibukalexander/file-processing/internal/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCalculator is a mock for the Calculator interface
type MockCalculator struct {
	mock.Mock
}

func (m *MockCalculator) Calculate(content string) (string, error) {
	args := m.Called(content)
	return args.String(0), args.Error(1)
}

func setupLoggingTest() (*MockCalculator, *bytes.Buffer) {
	config.AppConfig = &config.Config{EnableLogging: true}
	logger.SetupLogger()

	mockCalc := new(MockCalculator)
	logOutput := new(bytes.Buffer)
	logger.GetInstance().SetOutput(logOutput)
	// No defer to reset, as each test should control its logger state.

	return mockCalc, logOutput
}

func TestLoggingCalculator_Success(t *testing.T) {
	// Setup
	mockCalc, logOutput := setupLoggingTest()
	originalOutput := logger.GetInstance().Out
	defer logger.GetInstance().SetOutput(originalOutput)

	// Expectations
	input := "2 + 2"
	expectedResult := "4"
	mockCalc.On("Calculate", input).Return(expectedResult, nil)

	// Execution
	loggingCalc := NewLoggingCalculator(mockCalc)
	result, err := loggingCalc.Calculate(input)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Contains(t, logOutput.String(), "Starting calculation")
	assert.Contains(t, logOutput.String(), "Calculation finished")
	mockCalc.AssertExpectations(t)
}

func TestLoggingCalculator_Error(t *testing.T) {
	// Setup
	mockCalc, logOutput := setupLoggingTest()
	originalOutput := logger.GetInstance().Out
	defer logger.GetInstance().SetOutput(originalOutput)

	// Expectations
	input := "invalid"
	expectedError := errors.New("calculation error")
	mockCalc.On("Calculate", input).Return("", expectedError)

	// Execution
	loggingCalc := NewLoggingCalculator(mockCalc)
	_, err := loggingCalc.Calculate(input)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Contains(t, logOutput.String(), "Starting calculation")
	assert.Contains(t, logOutput.String(), "Calculation failed")
	mockCalc.AssertExpectations(t)
}
