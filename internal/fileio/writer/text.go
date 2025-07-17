package writer

import (
	"os"
	"strings"

	"github.com/dzibukalexander/file-processing/internal/fileio/constants"
)

type TextWriter struct{}

// '0644 FileMode' means you can read and write the file or
//
//	directory and other users can only read it.
//	Suitable for public text files.
func (w *TextWriter) Write(filePath string, data []byte) error {
	ext := strings.ToLower(string(constants.TEXT))
	basePath := strings.TrimSuffix(filePath, "."+ext)
	outputPath := basePath + "." + ext
	return os.WriteFile(outputPath, data, 0644)
}
