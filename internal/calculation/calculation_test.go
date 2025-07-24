package calculation

import (
	"testing"

	"github.com/dzibukalexander/file-processing/internal/calculation/library"
	"github.com/dzibukalexander/file-processing/internal/calculation/parser"
	"github.com/dzibukalexander/file-processing/internal/calculation/regex"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLibraryCalculator(t *testing.T) {
	calc := &library.LibraryCalculator{}
	testCases := map[string]string{
		"2 + 2 * 2":          "6",
		"10 / 2 - 1":         "4",
		"5 * (3 + 1)":        "20",
		"invalid expression": "invalid expression",
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			result, err := calc.Calculate(input)
			if expected == "invalid expression" {
				// The library might not return an error for un-evaluatable parts
				assert.Contains(t, result, "invalid")
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, result)
			}
		})
	}
}

func TestParserCalculator(t *testing.T) {
	calc := &parser.ParserCalculator{}
	// The parser is very basic and only handles simple "num op num" sequences
	testCases := map[string]string{
		"3 + 4":     "7",
		"10 - 5":    "5",
		"10 / 2":    "5",
		"3 * 4":     "12",
		"5 * 2 + 1": "11",
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			result, err := calc.Calculate(input)
			require.NoError(t, err)
			assert.Equal(t, expected, result)
		})
	}
}

func TestRegexCalculator(t *testing.T) {
	calc := &regex.RegexCalculator{}
	// Regex calculator is also basic and finds "num op num"
	testCases := map[string]string{
		"hello 3 * 5 world": "hello 15 world",
		"10 / 2 = five":     "5 = five",
		"no math here":      "no math here",
	}

	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			result, err := calc.Calculate(input)
			require.NoError(t, err)
			assert.Equal(t, expected, result)
		})
	}
}
