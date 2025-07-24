// internal/compression/interface.go
package compression

import (
	. "github.com/dzibukalexander/file-processing/internal/compression/constants"
	"github.com/dzibukalexander/file-processing/internal/compression/gzip"
	"github.com/dzibukalexander/file-processing/internal/compression/zip"
)

type Compressor interface {
	Compress(data []byte) ([]byte, error)
}

type Decompressor interface {
	Decompress(data []byte) ([]byte, error)
}

func NewCompressor(compType CompressionType) Compressor {
	var compressor Compressor
	switch compType {
	case GZIP:
		compressor = &gzip.GzipCompressor{}
	case ZIP:
		compressor = &zip.ZipCompressor{}
	default:
		return nil
	}
	return NewLoggingCompressor(compressor)
}

func NewDecompressor(compType CompressionType) Decompressor {
	var decompressor Decompressor
	switch compType {
	case GZIP:
		decompressor = &gzip.GzipDecompressor{}
	case ZIP:
		decompressor = &zip.ZipDecompressor{}
	default:
		return nil
	}
	return NewLoggingDecompressor(decompressor)
}
