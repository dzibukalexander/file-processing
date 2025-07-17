package writer

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/dzibukalexander/file-processing/internal/fileio/constants"
)

type JSONWriter struct{}

// '0644 FileMode' means you can read and write the file or
//
//	directory and other users can only read it.
//	Suitable for public text files.
func (w *JSONWriter) Write(filePath string, data []byte) error {
	// Validate data is valid JSON
	var temp interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	ext := strings.ToLower(string(constants.JSON))
	basePath := strings.TrimSuffix(filePath, "."+ext)
	outputPath := basePath + "." + ext

	return os.WriteFile(outputPath, data, 0644)
}
