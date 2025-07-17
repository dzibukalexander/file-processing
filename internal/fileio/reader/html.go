package reader

import "os"

type HTMLReader struct{}

func (h *HTMLReader) Read(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
