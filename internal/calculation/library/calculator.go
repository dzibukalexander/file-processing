package library

import (
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

type LibraryCalculator struct{}

func (c *LibraryCalculator) Calculate(content string) (string, error) {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		expression, err := govaluate.NewEvaluableExpression(line)
		if err != nil {
			continue
		}

		result, err := expression.Evaluate(nil)
		if err == nil {
			switch v := result.(type) {
			case float64:
				lines[i] = strings.TrimRight(strings.TrimRight(strconv.FormatFloat(v, 'f', 6, 64), "0"), ".")
			case int:
				lines[i] = strconv.Itoa(v)
			}
		}
	}
	return strings.Join(lines, "\n"), nil
}
