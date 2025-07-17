// internal/calculator/parser/calculator.go
package parser

import (
	"strconv"
	"strings"

	"github.com/dzibukalexander/file-processing/internal/arithmetic"
)

type ParserCalculator struct{}

func (c *ParserCalculator) Calculate(content string) (string, error) {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		tokens := strings.Fields(line)
		for j := 0; j < len(tokens)-2; j++ {
			if arithmetic.IsNumber(tokens[j]) &&
				arithmetic.IsOperator(tokens[j+1]) &&
				arithmetic.IsNumber(tokens[j+2]) {
				a, _ := strconv.Atoi(tokens[j])
				b, _ := strconv.Atoi(tokens[j+2])
				op, err := arithmetic.OperatorFromString(tokens[j+1])
				if err != nil {
					continue
				}

				res, err := arithmetic.ApplyOperator(a, b, op)
				if err != nil {
					continue
				}

				tokens[j] = arithmetic.FormatResult(res)
				tokens = append(tokens[:j+1], tokens[j+3:]...)
				j--
			}
		}
		lines[i] = strings.Join(tokens, " ")
	}
	return strings.Join(lines, "\n"), nil
}
