package calculation

import (
	"testing"

	"github.com/dzibukalexander/file-processing/internal/calculation/library"
	"github.com/dzibukalexander/file-processing/internal/calculation/parser"
	"github.com/dzibukalexander/file-processing/internal/calculation/regex"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestLibraryCalculator(t *testing.T) {
	calc := &library.LibraryCalculator{}
	testCases := map[string]string{
		"2 + 2 * 2":          "6",
		"10 / 2 - 1":         "4",
		"5 * (3 + 1)":        "20",
		"invalid expression": "invalid expression",
	}

	runner.Run(t, "LibraryCalculator", func(t provider.T) {
		for input, expected := range testCases {
			in, exp := input, expected
			t.WithNewStep(in, func(s provider.StepCtx) {
				result, err := calc.Calculate(in)
				if exp == "invalid expression" {
					s.Assert().Contains(result, "invalid")
				} else {
					s.Require().NoError(err)
					s.Assert().Equal(exp, result)
				}
			})
		}
	})
}

func TestParserCalculator(t *testing.T) {
	calc := &parser.ParserCalculator{}
	testCases := map[string]string{
		"3 + 4":                   "7",
		"10 - 5":                  "5",
		"10 / 2":                  "5",
		"3 * 4":                   "12",
		"5 * 2 + 1":               "11",
		"2 + 2 * 2":               "6",
		"(2 + 2) * 2":             "8",
		"(5*(5+3)*(2*(3-5)))":     "-160",
		"1 + (2 + (3 + (4 + 5)))": "15",
		"(((((7)))))":             "7",
		"2 * (3 + 4 * (2 + 1))":   "30",
		"(1 + 2) * (3 + 4)":       "21",
		"(8 / (4 - 2)) + 1":       "5",
		"(2 + 3) * (4 + 5)":       "45",
		"42":                      "42",
		"((1+2)*(3+4))/7":         "3",
	}

	runner.Run(t, "ParserCalculator", func(t provider.T) {
		for input, expected := range testCases {
			in, exp := input, expected
			t.WithNewStep(in, func(s provider.StepCtx) {
				result, err := calc.Calculate(in)
				s.Require().NoError(err)
				s.Assert().Equal(exp, result)
			})
		}
	})
}

func TestRegexCalculator(t *testing.T) {
	calc := &regex.RegexCalculator{}
	testCases := map[string]string{
		"hello 3 * 5 world": "hello 15 world",
		"10 / 2 = five":     "5 = five",
		"no math here":      "no math here",
	}

	runner.Run(t, "RegexCalculator", func(t provider.T) {
		for input, expected := range testCases {
			in, exp := input, expected
			t.WithNewStep(in, func(s provider.StepCtx) {
				result, err := calc.Calculate(in)
				s.Require().NoError(err)
				s.Assert().Equal(exp, result)
			})
		}
	})
}
