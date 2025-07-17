package writer

import (
	"os"
	"strings"

	"github.com/dzibukalexander/file-processing/internal/fileio/constants"
)

type XMLWriter struct{}

// '0644 FileMode' means you can read and write the file or
//
//	directory and other users can only read it.
//	Suitable for public text files.
func (x *XMLWriter) Write(filePath string, data []byte) error {
	ext := strings.ToLower(string(constants.XML))
	basePath := strings.TrimSuffix(filePath, "."+ext)
	outputPath := basePath + "." + ext
	return os.WriteFile(outputPath, data, 0644)
}
