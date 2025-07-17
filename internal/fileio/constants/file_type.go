package constants

import "fmt"

type FileType string

const (
	TEXT FileType = "TEXT"
	JSON FileType = "JSON"
	XML  FileType = "XML"
	YAML FileType = "YAML"
	HTML FileType = "HTML"
)

func FileTypeFromString(s string) (FileType, error) {
	switch s {
	case "TEXT":
		return TEXT, nil
	case "JSON":
		return JSON, nil
	case "XML":
		return XML, nil
	case "YAML":
		return YAML, nil
	case "HTML":
		return HTML, nil
	default:
		return "", fmt.Errorf("unknown file type: %s", s)
	}
}
