package reader

import (
	"encoding/json"
	"os"
)

type JSONReader struct{}

func (r *JSONReader) Read(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var temp interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return nil, err
	}

	return data, nil
}
