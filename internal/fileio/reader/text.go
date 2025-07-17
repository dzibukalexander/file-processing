package reader

import "os"

type TextReader struct{}

func (r *TextReader) Read(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
