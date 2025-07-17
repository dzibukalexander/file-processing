package constants

import "fmt"

type ArithmeticOperator string

const (
	ADD      ArithmeticOperator = "+"
	SUBTRACT ArithmeticOperator = "-"
	MULTIPLY ArithmeticOperator = "*"
	DIVIDE   ArithmeticOperator = "/"
)

func (op ArithmeticOperator) String() string {
	return string(op)
}

func OperatorFromString(s string) (ArithmeticOperator, error) {
	switch s {
	case "+":
		return ADD, nil
	case "-":
		return SUBTRACT, nil
	case "*":
		return MULTIPLY, nil
	case "/":
		return DIVIDE, nil
	default:
		return "", fmt.Errorf("unknown operator: %s", s)
	}
}
