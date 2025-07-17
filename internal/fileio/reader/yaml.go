package reader

import "os"

type YAMLReader struct{}

func (y *YAMLReader) Read(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
