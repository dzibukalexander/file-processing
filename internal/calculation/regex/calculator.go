package regex

import (
	"regexp"
	"strconv"
)

type RegexCalculator struct{}

func (c *RegexCalculator) Calculate(content string) (string, error) {
	re := regexp.MustCompile(`\b(\d+)\s*([+\-*\/])\s*(\d+)\b`)
	result := re.ReplaceAllStringFunc(content, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) != 4 {
			return match
		}

		a, _ := strconv.Atoi(parts[1])
		b, _ := strconv.Atoi(parts[3])
		op := parts[2]

		var res int
		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			if b == 0 {
				return match
			}
			res = a / b
		}
		return strconv.Itoa(res)
	})
	return result, nil
}
