package constants

import "fmt"

type EncryptionType string

const (
	NONE EncryptionType = "NONE"
	AES  EncryptionType = "AES"
	RSA  EncryptionType = "RSA"
)

func EncryptionTypeFromString(s string) (EncryptionType, error) {
	switch s {
	case "NONE":
		return NONE, nil
	case "AES":
		return AES, nil
	case "RSA":
		return RSA, nil
	default:
		return NONE, fmt.Errorf("unknown encryption type: %s", s)
	}
}
