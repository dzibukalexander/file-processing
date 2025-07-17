// fileio.go
package fileio

import (
	"github.com/dzibukalexander/file-processing/internal/fileio/constants"
	"github.com/dzibukalexander/file-processing/internal/fileio/reader"
	"github.com/dzibukalexander/file-processing/internal/fileio/writer"
)

type FileReader interface {
	Read(filePath string) ([]byte, error)
}

type FileWriter interface {
	Write(filePath string, data []byte) error
}

type FileIO interface {
	FileReader
	FileWriter
}

func NewFileReader(fileType constants.FileType) FileReader {
	switch fileType {
	case constants.TEXT:
		return &reader.TextReader{}
	case constants.JSON:
		return &reader.JSONReader{}
	case constants.YAML:
		return &reader.YAMLReader{}
	case constants.HTML:
		return &reader.HTMLReader{}
	default:
		return &reader.TextReader{}
	}
}

func NewWriter(fileType constants.FileType) FileWriter {
	switch fileType {
	case constants.JSON:
		return &writer.JSONWriter{}
	case constants.XML:
		return &writer.XMLWriter{}
	case constants.YAML:
		return &writer.YAMLWriter{}
	case constants.HTML:
		return &writer.HTMLWriter{}
	default:
		return &writer.TextWriter{}
	}
}
