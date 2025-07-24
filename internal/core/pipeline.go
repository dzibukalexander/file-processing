package core

import (
	"encoding/json"
	"io/ioutil"
)

// PipelineBuilder constructs a sequence of processing steps.
type PipelineBuilder struct {
	operations []*Operation
}

// NewPipelineBuilder creates a new builder.
func NewPipelineBuilder() *PipelineBuilder {
	return &PipelineBuilder{}
}

// Add appends a new operation to the pipeline.
func (b *PipelineBuilder) Add(op *Operation) {
	b.operations = append(b.operations, op)
}

// Reset clears all operations from the pipeline.
func (b *PipelineBuilder) Reset() {
	b.operations = []*Operation{}
}

// SaveToFile serializes the pipeline to a JSON file.
func (b *PipelineBuilder) SaveToFile(filePath string) error {
	data, err := json.MarshalIndent(b.operations, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

// LoadFromFile deserializes the pipeline from a JSON file.
func (b *PipelineBuilder) LoadFromFile(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &b.operations)
}
