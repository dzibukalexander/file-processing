package constants

import "fmt"

type CalculationMethod string

const (
	REGEX   CalculationMethod = "REGEX"
	PARSER  CalculationMethod = "PARSER"
	LIBRARY CalculationMethod = "LIBRARY"
)

func CalculationMethodFromString(s string) (CalculationMethod, error) {
	switch s {
	case "REGEX":
		return REGEX, nil
	case "PARSER":
		return PARSER, nil
	case "LIBRARY":
		return LIBRARY, nil
	default:
		return "", fmt.Errorf("unknown calculation method: %s", s)
	}
}
