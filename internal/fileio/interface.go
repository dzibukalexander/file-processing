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
	var r FileReader
	switch fileType {
	case constants.TEXT:
		r = &reader.TextReader{}
	case constants.JSON:
		r = &reader.JSONReader{}
	case constants.YAML:
		r = &reader.YAMLReader{}
	case constants.HTML:
		r = &reader.HTMLReader{}
	default:
		r = &reader.TextReader{}
	}
	return NewLoggingFileReader(r)
}

func NewWriter(fileType constants.FileType) FileWriter {
	var w FileWriter
	switch fileType {
	case constants.JSON:
		w = &writer.JSONWriter{}
	case constants.XML:
		w = &writer.XMLWriter{}
	case constants.YAML:
		w = &writer.YAMLWriter{}
	case constants.HTML:
		w = &writer.HTMLWriter{}
	default:
		w = &writer.TextWriter{}
	}
	return NewLoggingFileWriter(w)
}
