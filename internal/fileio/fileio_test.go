package fileio

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dzibukalexander/file-processing/internal/fileio/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) (string, func()) {
	tempDir, err := ioutil.TempDir("", "fileio-test")
	require.NoError(t, err)
	cleanup := func() { os.RemoveAll(tempDir) }
	return tempDir, cleanup
}

func TestFileIO_ReadWrite(t *testing.T) {
	testCases := []struct {
		fileType constants.FileType
		data     []byte
		filename string
	}{
		{constants.TEXT, []byte("hello text"), "test.txt"},
		{constants.JSON, []byte(`{"hello":"json"}`), "test.json"},
		{constants.XML, []byte(`<hello>xml</hello>`), "test.xml"},
		{constants.YAML, []byte(`hello: yaml`), "test.yaml"},
		{constants.HTML, []byte(`<h1>hello html</h1>`), "test.html"},
	}

	for _, tc := range testCases {
		t.Run(string(tc.fileType), func(t *testing.T) {
			tempDir, cleanup := setupTest(t)
			defer cleanup()

			filePath := filepath.Join(tempDir, tc.filename)

			writer := NewWriter(tc.fileType)
			err := writer.Write(filePath, tc.data)
			require.NoError(t, err)

			reader := NewFileReader(tc.fileType)
			readData, err := reader.Read(filePath)
			require.NoError(t, err)

			// YAML and JSON writers/readers might slightly reformat, so we do a content-wise check
			// For simplicity here, we'll stick to direct byte comparison.
			// In a real-world scenario, you might unmarshal and compare the structures.
			assert.Equal(t, tc.data, readData)
		})
	}
}
