package arithmetic

import (
	"errors"
	"fmt"
)

func ApplyOperator(a int, b int, op Operator) (float64, error) {
	switch op {
	case Add:
		return float64(a + b), nil
	case Subtract:
		return float64(a - b), nil
	case Multiply:
		return float64(a * b), nil
	case Divide:
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return float64(a) / float64(b), nil
	default:
		return 0, fmt.Errorf("unknown operator: %v", op)
	}
}
