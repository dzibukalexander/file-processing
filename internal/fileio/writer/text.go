package writer

import (
	"os"
)

type TextWriter struct{}

// '0644 FileMode' means you can read and write the file or
//
//	directory and other users can only read it.
//	Suitable for public text files.
func (w *TextWriter) Write(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}
