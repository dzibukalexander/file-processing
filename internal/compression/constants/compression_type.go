// internal/compression/constants.go
package compression

import "fmt"

type CompressionType string

const (
	NONE CompressionType = "NONE"
	GZIP CompressionType = "GZIP"
	ZIP  CompressionType = "ZIP"
)

func CompressionTypeFromString(s string) (CompressionType, error) {
	switch s {
	case "NONE":
		return NONE, nil
	case "GZIP":
		return GZIP, nil
	case "ZIP":
		return ZIP, nil
	default:
		return NONE, fmt.Errorf("unknown compression type: %s", s)
	}
}
