// internal/calculator/parser/calculator.go
package parser

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/dzibukalexander/file-processing/internal/arithmetic"
)

type ParserCalculator struct{}

// precedence and associativity
var precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func isOperator(tok string) bool {
	return tok == "+" || tok == "-" || tok == "*" || tok == "/"
}

func shuntingYard(tokens []string) ([]string, error) {
	var output []string
	var stack []string
	for _, tok := range tokens {
		switch {
		case IsNumber(tok):
			output = append(output, tok)
		case isOperator(tok):
			for len(stack) > 0 && isOperator(stack[len(stack)-1]) &&
				(precedence[tok] <= precedence[stack[len(stack)-1]]) {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, tok)
		case tok == "(":
			stack = append(stack, tok)
		case tok == ")":
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			stack = stack[:len(stack)-1] // pop "("
		default:
			return nil, fmt.Errorf("invalid token: %s", tok)
		}
	}
	for len(stack) > 0 {
		if stack[len(stack)-1] == "(" || stack[len(stack)-1] == ")" {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return output, nil
}

func evalRPN(rpn []string) (int, error) {
	var stack []int
	for _, tok := range rpn {
		if IsNumber(tok) {
			n, _ := strconv.Atoi(tok)
			stack = append(stack, n)
		} else if isOperator(tok) {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			op, err := OperatorFromString(tok)
			if err != nil {
				return 0, err
			}
			res, err := ApplyOperator(a, b, op)
			if err != nil {
				return 0, err
			}
			stack = append(stack, int(res))
		} else {
			return 0, fmt.Errorf("invalid token in RPN: %s", tok)
		}
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}
	return stack[0], nil
}

func tokenize(expr string) []string {
	var tokens []string
	runes := []rune(expr)
	n := len(runes)
	i := 0
	for i < n {
		if runes[i] == ' ' {
			i++
			continue
		}
		if runes[i] >= '0' && runes[i] <= '9' {
			j := i
			for j < n && runes[j] >= '0' && runes[j] <= '9' {
				j++
			}
			tokens = append(tokens, string(runes[i:j]))
			i = j
		} else if strings.ContainsRune("+-*/()", runes[i]) {
			tokens = append(tokens, string(runes[i]))
			i++
		} else {
			tokens = append(tokens, string(runes[i]))
			i++
		}
	}
	return tokens
}

func (c *ParserCalculator) Calculate(content string) (string, error) {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		tokens := tokenize(line)
		if len(tokens) == 0 {
			continue
		}
		rpn, err := shuntingYard(tokens)
		if err != nil {
			lines[i] = line
			continue
		}
		res, err := evalRPN(rpn)
		if err != nil {
			lines[i] = line
			continue
		}
		lines[i] = strconv.Itoa(res)
	}
	return strings.Join(lines, "\n"), nil
}
