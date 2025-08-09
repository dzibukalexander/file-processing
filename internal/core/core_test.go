package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func setupTest(t *testing.T) (string, func()) {
	tempDir, err := ioutil.TempDir("", "core-test")
	if err != nil {
		t.Fatalf("temp dir: %v", err)
	}

	cleanup := func() {
		os.RemoveAll(tempDir)
	}

	return tempDir, cleanup
}

func TestCore_Load(t *testing.T) {
	tempDir, cleanup := setupTest(t)
	defer cleanup()

	filePath := filepath.Join(tempDir, "input.txt")
	if err := ioutil.WriteFile(filePath, []byte("hello"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	runner.Run(t, "Core Load", func(t provider.T) {
		t.WithNewStep("load file", func(s provider.StepCtx) {
			core := NewCore()
			err := core.Load(filePath)
			s.Assert().NoError(err)
			s.Assert().Equal([]byte("hello"), core.originalData)
		})
	})
}

func TestCore_Apply(t *testing.T) {
	runner.Run(t, "Core Apply", func(t provider.T) {
		t.WithNewStep("add step", func(s provider.StepCtx) {
			core := NewCore()
			params := map[string]string{"type": "zip"}
			err := core.Apply("compress", params)
			s.Require().NoError(err)
			s.Assert().Len(core.builder.operations, 1)
			s.Assert().Equal("compress", core.builder.operations[0].Name)
		})
	})
}

func TestCore_ProcessFile(t *testing.T) {
	tempDir, cleanup := setupTest(t)
	defer cleanup()

	inputPath := filepath.Join(tempDir, "input.txt")
	if err := ioutil.WriteFile(inputPath, []byte("1 + 1"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	runner.Run(t, "Core ProcessFile", func(t provider.T) {
		t.WithNewStep("process", func(s provider.StepCtx) {
			core := NewCore()
			s.Require().NoError(core.Load(inputPath))
			s.Require().NoError(core.Apply("calculate", map[string]string{"type": "library"}))

			outputPath := filepath.Join(tempDir, "output.txt")
			s.Require().NoError(core.ProcessFile(outputPath))

			outputData, err := ioutil.ReadFile(outputPath)
			s.Require().NoError(err)
			s.Assert().Equal("2", string(outputData))
		})
	})
}

func TestCore_SaveLoadPipeline(t *testing.T) {
	tempDir, cleanup := setupTest(t)
	defer cleanup()

	runner.Run(t, "Core Save/Load pipeline", func(t provider.T) {
		t.WithNewStep("save and load", func(s provider.StepCtx) {
			core := NewCore()
			s.Require().NoError(core.Apply("encrypt", map[string]string{"type": "aes", "key_file": "dummy.key"}))
			s.Require().NoError(core.Apply("compress", map[string]string{"type": "gzip"}))

			pipelinePath := filepath.Join(tempDir, "pipeline.json")
			s.Require().NoError(core.SavePipeline(pipelinePath))

			newCore := NewCore()
			s.Require().NoError(newCore.LoadPipeline(pipelinePath))
			s.Assert().Len(newCore.builder.operations, 2)
			s.Assert().Equal("encrypt", newCore.builder.operations[0].Name)
			s.Assert().Equal("compress", newCore.builder.operations[1].Name)
		})
	})
}
