package arithmetic

import "fmt"

type Operator string

const (
	Add      Operator = "+"
	Subtract Operator = "-"
	Multiply Operator = "*"
	Divide   Operator = "/"
)

func OperatorFromString(s string) (Operator, error) {
	switch s {
	case "+":
		return Add, nil
	case "-":
		return Subtract, nil
	case "*":
		return Multiply, nil
	case "/":
		return Divide, nil
	default:
		return "", fmt.Errorf("unknown operator: %s", s)
	}
}

func IsOperator(s string) bool {
	_, err := OperatorFromString(s)
	return err == nil
}
