package constants

import (
	"fmt"
	"path/filepath"
	"strings"
)

type FileType string

const (
	TEXT FileType = "TEXT"
	JSON FileType = "JSON"
	XML  FileType = "XML"
	YAML FileType = "YAML"
	HTML FileType = "HTML"
)

func FileTypeFromExtension(filename string) (FileType, error) {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".txt", ".text":
		return TEXT, nil
	case ".json":
		return JSON, nil
	case ".xml":
		return XML, nil
	case ".yaml", ".yml":
		return YAML, nil
	case ".html", ".htm":
		return HTML, nil
	default:
		return "", fmt.Errorf("unknown file extension: %s", ext)
	}
}

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
