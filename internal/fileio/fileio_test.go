package fileio

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dzibukalexander/file-processing/internal/fileio/constants"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func setupTest(t *testing.T) (string, func()) {
	tempDir, err := ioutil.TempDir("", "fileio-test")
	if err != nil {
		t.Fatalf("temp dir: %v", err)
	}
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

	runner.Run(t, "FileIO Read/Write", func(at provider.T) {
		for _, tc := range testCases {
			tc := tc
			tempDir, cleanup := setupTest(t)
			defer cleanup()

			at.WithNewStep(string(tc.fileType), func(s provider.StepCtx) {
				filePath := filepath.Join(tempDir, tc.filename)

				writer := NewWriter(tc.fileType)
				s.Require().NoError(writer.Write(filePath, tc.data))

				reader := NewFileReader(tc.fileType)
				readData, err := reader.Read(filePath)
				s.Require().NoError(err)
				s.Assert().Equal(tc.data, readData)
			})
		}
	})
}
