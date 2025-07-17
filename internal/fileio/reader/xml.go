package reader

import "os"

type XMLReader struct{}

func (x *XMLReader) Read(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
