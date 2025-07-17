package writer

import (
	"os"
	"strings"

	"github.com/dzibukalexander/file-processing/internal/fileio/constants"
)

type YAMLWriter struct{}

// '0644 FileMode' means you can read and write the file or
//
//	directory and other users can only read it.
//	Suitable for public text files.
func (y *YAMLWriter) Write(filePath string, data []byte) error {
	ext := strings.ToLower(string(constants.YAML))
	basePath := strings.TrimSuffix(filePath, "."+ext)
	outputPath := basePath + "." + ext
	return os.WriteFile(outputPath, data, 0644)
}
