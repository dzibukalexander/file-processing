package compression

import (
	"testing"

	"github.com/dzibukalexander/file-processing/internal/compression/gzip"
	"github.com/dzibukalexander/file-processing/internal/compression/zip"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestGzipCompressDecompress_String(t *testing.T) {
	runner.Run(t, "Gzip string roundtrip", func(t provider.T) {
		t.WithNewStep("roundtrip", func(s provider.StepCtx) {
			compressor := &gzip.GzipCompressor{}
			decompressor := &gzip.GzipDecompressor{}
			original := "hello gzip"
			originalData := []byte(original)

			compressed, err := compressor.Compress(originalData)
			s.Require().NoError(err)
			s.Assert().NotEqual(originalData, compressed)

			decompressed, err := decompressor.Decompress(compressed)
			s.Require().NoError(err)
			decompressedStr := string(decompressed)
			s.Assert().Equal(original, decompressedStr)
		})
	})
}

func TestZipCompressDecompress_String(t *testing.T) {
	runner.Run(t, "Zip string roundtrip", func(t provider.T) {
		t.WithNewStep("roundtrip", func(s provider.StepCtx) {
			compressor := &zip.ZipCompressor{}
			decompressor := &zip.ZipDecompressor{}
			original := "hello zip"
			originalData := []byte(original)

			compressed, err := compressor.Compress(originalData)
			s.Require().NoError(err)
			s.Assert().NotEqual(originalData, compressed)

			decompressed, err := decompressor.Decompress(compressed)
			s.Require().NoError(err)
			decompressedStr := string(decompressed)
			s.Assert().Equal(original, decompressedStr)
		})
	})
}
