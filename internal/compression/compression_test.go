package compression

import (
	"testing"

	"github.com/dzibukalexander/file-processing/internal/compression/gzip"
	"github.com/dzibukalexander/file-processing/internal/compression/zip"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGzipCompressDecompress(t *testing.T) {
	compressor := &gzip.GzipCompressor{}
	decompressor := &gzip.GzipDecompressor{}
	originalData := []byte("hello gzip")

	compressed, err := compressor.Compress(originalData)
	require.NoError(t, err)

	decompressed, err := decompressor.Decompress(compressed)
	require.NoError(t, err)

	assert.Equal(t, originalData, decompressed)
}

func TestZipCompressDecompress(t *testing.T) {
	compressor := &zip.ZipCompressor{}
	decompressor := &zip.ZipDecompressor{}
	originalData := []byte("hello zip")

	compressed, err := compressor.Compress(originalData)
	require.NoError(t, err)

	decompressed, err := decompressor.Decompress(compressed)
	require.NoError(t, err)

	assert.Equal(t, originalData, decompressed)
}
