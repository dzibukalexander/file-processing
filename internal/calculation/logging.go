package calculation

import (
	"time"

	"github.com/dzibukalexander/file-processing/internal/logger"
)

type loggingCalculator struct {
	calculator Calculator
}

func NewLoggingCalculator(calculator Calculator) Calculator {
	return &loggingCalculator{calculator: calculator}
}

func (l *loggingCalculator) Calculate(content string) (result string, err error) {
	log := logger.GetInstance().WithField("input_size", len(content))
	log.Info("Starting calculation")

	defer func(begin time.Time) {
		if err != nil {
			log.WithError(err).Error("Calculation failed")
		} else {
			log.WithField("duration", time.Since(begin)).Info("Calculation finished")
		}
	}(time.Now())

	return l.calculator.Calculate(content)
}
