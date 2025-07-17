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
	switch compType {
	case GZIP:
		return &gzip.GzipCompressor{}
	case ZIP:
		return &zip.ZipCompressor{}
	default:
		return nil
	}
}

func NewDecompressor(compType CompressionType) Decompressor {
	switch compType {
	case GZIP:
		return &gzip.GzipDecompressor{}
	case ZIP:
		return &zip.ZipDecompressor{}
	default:
		return nil
	}
}
