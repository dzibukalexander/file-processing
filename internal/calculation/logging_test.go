package calculation

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
	return mockCalc, logOutput
}

func TestLoggingCalculator_Success(t *testing.T) {
	runner.Run(t, "LoggingCalculator success", func(t provider.T) {
		t.WithNewStep("success path", func(s provider.StepCtx) {
			mockCalc, logOutput := setupLoggingTest()
			origOut := logger.GetInstance().Out
			defer logger.GetInstance().SetOutput(origOut)

			input := "2 + 2"
			expectedResult := "4"
			mockCalc.On("Calculate", input).Return(expectedResult, nil)

			loggingCalc := NewLoggingCalculator(mockCalc)
			result, err := loggingCalc.Calculate(input)

			s.Assert().NoError(err)
			s.Assert().Equal(expectedResult, result)
			s.Assert().Contains(logOutput.String(), "Starting calculation")
			s.Assert().Contains(logOutput.String(), "Calculation finished")
			mockCalc.AssertExpectations(t)
		})
	})
}

func TestLoggingCalculator_Error(t *testing.T) {
	runner.Run(t, "LoggingCalculator error", func(t provider.T) {
		t.WithNewStep("error path", func(s provider.StepCtx) {
			mockCalc, logOutput := setupLoggingTest()
			origOut := logger.GetInstance().Out
			defer logger.GetInstance().SetOutput(origOut)

			input := "invalid"
			expectedError := errors.New("calculation error")
			mockCalc.On("Calculate", input).Return("", expectedError)

			loggingCalc := NewLoggingCalculator(mockCalc)
			_, err := loggingCalc.Calculate(input)

			s.Assert().Error(err)
			s.Assert().Equal(expectedError, err)
			s.Assert().Contains(logOutput.String(), "Starting calculation")
			s.Assert().Contains(logOutput.String(), "Calculation failed")
			mockCalc.AssertExpectations(t)
		})
	})
}
