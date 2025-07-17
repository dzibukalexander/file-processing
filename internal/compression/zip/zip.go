// internal/compression/zip/zip.go
package zip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
)

type ZipCompressor struct{}

func (c *ZipCompressor) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	f, err := w.Create("data")
	if err != nil {
		return nil, err
	}
	if _, err := f.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type ZipDecompressor struct{}

func (d *ZipDecompressor) Decompress(data []byte) ([]byte, error) {
	r, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}
	if len(r.File) == 0 {
		return nil, fmt.Errorf("no files in zip archive")
	}
	rc, err := r.File[0].Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return io.ReadAll(rc)
}
