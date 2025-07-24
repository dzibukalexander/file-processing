package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) (string, func()) {
	tempDir, err := ioutil.TempDir("", "core-test")
	require.NoError(t, err)

	cleanup := func() {
		os.RemoveAll(tempDir)
	}

	return tempDir, cleanup
}

func TestCore_Load(t *testing.T) {
	tempDir, cleanup := setupTest(t)
	defer cleanup()

	filePath := filepath.Join(tempDir, "input.txt")
	err := ioutil.WriteFile(filePath, []byte("hello"), 0644)
	require.NoError(t, err)

	core := NewCore()
	err = core.Load(filePath)
	assert.NoError(t, err)
	assert.Equal(t, []byte("hello"), core.originalData)
}

func TestCore_Apply(t *testing.T) {
	core := NewCore()
	params := map[string]string{"type": "zip"}
	err := core.Apply("compress", params)

	require.NoError(t, err)
	assert.Len(t, core.builder.operations, 1)
	assert.Equal(t, "compress", core.builder.operations[0].Name)
}

func TestCore_ProcessFile(t *testing.T) {
	tempDir, cleanup := setupTest(t)
	defer cleanup()

	// Setup input file
	inputPath := filepath.Join(tempDir, "input.txt")
	err := ioutil.WriteFile(inputPath, []byte("1 + 1"), 0644)
	require.NoError(t, err)

	// Setup core and pipeline
	core := NewCore()
	err = core.Load(inputPath)
	require.NoError(t, err)
	err = core.Apply("calculate", map[string]string{"type": "library"})
	require.NoError(t, err)

	// Process and check output
	outputPath := filepath.Join(tempDir, "output.txt")
	err = core.ProcessFile(outputPath)
	require.NoError(t, err)

	outputData, err := ioutil.ReadFile(outputPath)
	require.NoError(t, err)
	assert.Equal(t, "2", string(outputData))
}

func TestCore_SaveLoadPipeline(t *testing.T) {
	tempDir, cleanup := setupTest(t)
	defer cleanup()

	core := NewCore()
	err := core.Apply("encrypt", map[string]string{"type": "aes", "key_file": "dummy.key"})
	require.NoError(t, err)
	err = core.Apply("compress", map[string]string{"type": "gzip"})
	require.NoError(t, err)

	pipelinePath := filepath.Join(tempDir, "pipeline.json")
	err = core.SavePipeline(pipelinePath)
	require.NoError(t, err)

	newCore := NewCore()
	err = newCore.LoadPipeline(pipelinePath)
	require.NoError(t, err)
	assert.Len(t, newCore.builder.operations, 2)
	assert.Equal(t, "encrypt", newCore.builder.operations[0].Name)
	assert.Equal(t, "compress", newCore.builder.operations[1].Name)
}
