package calculation

import (
	"github.com/dzibukalexander/file-processing/internal/calculation/constants"
	"github.com/dzibukalexander/file-processing/internal/calculation/library"
	"github.com/dzibukalexander/file-processing/internal/calculation/parser"
	"github.com/dzibukalexander/file-processing/internal/calculation/regex"
)

type Calculator interface {
	Calculate(content string) (string, error)
}

func NewCalculator(method constants.CalculationMethod) Calculator {
	var calculator Calculator
	switch method {
	case constants.PARSER:
		calculator = &parser.ParserCalculator{}
	case constants.LIBRARY:
		calculator = &library.LibraryCalculator{}
	default:
		calculator = &regex.RegexCalculator{}
	}
	return NewLoggingCalculator(calculator)
}
