package core

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/dzibukalexander/file-processing/internal/calculation"
	calc_const "github.com/dzibukalexander/file-processing/internal/calculation/constants"
	"github.com/dzibukalexander/file-processing/internal/compression"
	comp_const "github.com/dzibukalexander/file-processing/internal/compression/constants"
	"github.com/dzibukalexander/file-processing/internal/encryption"
	enc_const "github.com/dzibukalexander/file-processing/internal/encryption/constants"
	"github.com/dzibukalexander/file-processing/internal/fileio"
	"github.com/dzibukalexander/file-processing/internal/fileio/constants"
	"github.com/dzibukalexander/file-processing/internal/logger"
)

// Core is the central part of the application, managing data and the processing pipeline.
type Core struct {
	originalData []byte
	builder      *PipelineBuilder
}

// NewCore creates a new Core instance.
func NewCore() *Core {
	return &Core{
		builder: NewPipelineBuilder(),
	}
}

// Load reads a file into memory and resets the processing pipeline.
func (c *Core) Load(filePath string) error {
	log := logger.GetInstance()
	c.builder.Reset()
	log.Debug("Pipeline builder reset")
	fileType, err := constants.FileTypeFromExtension(filePath)
	if err != nil {
		log.WithField("path", filePath).Errorf("Failed to determine file type: %v", err)
		return fmt.Errorf("failed to determine file type: %w", err)
	}

	reader := fileio.NewFileReader(fileType)
	data, err := reader.Read(filePath)
	if err != nil {
		log.WithField("path", filePath).Errorf("Failed to read file: %v", err)
		return fmt.Errorf("failed to read file: %w", err)
	}

	c.originalData = data
	log.WithFields(map[string]interface{}{
		"path": filePath,
		"size": len(data),
	}).Info("File loaded successfully")
	return nil
}

// ProcessFile builds and runs the pipeline, then writes the result to a file.
func (c *Core) ProcessFile(filePath string) error {
	log := logger.GetInstance()
	if c.originalData == nil {
		log.Warn("ProcessFile called with no data loaded")
		return fmt.Errorf("no data loaded to process")
	}
	log.Info("Starting file processing pipeline")

	data := c.originalData
	var err error

	type Step func([]byte) ([]byte, error)

	for i, op := range c.builder.operations {
		log.WithFields(map[string]interface{}{
			"step":      i + 1,
			"operation": op.Name,
			"params":    op.Params,
		}).Debug("Executing pipeline step")
		step, err_step := c.createStep(op.Name, op.Params)
		if err_step != nil {
			log.Errorf("Error creating step %d (%s): %v", i+1, op.Name, err_step)
			return err_step
		}
		data, err = step(data)
		if err != nil {
			log.Errorf("Error processing step %d (%s): %v", i+1, op.Name, err)
			return fmt.Errorf("error processing step '%s': %w", op.Name, err)
		}
	}

	outputType, _ := constants.FileTypeFromExtension(filePath)
	writer := fileio.NewWriter(outputType)
	if err := writer.Write(filePath, data); err != nil {
		log.WithField("path", filePath).Errorf("Failed to write file: %v", err)
		return fmt.Errorf("failed to write file: %w", err)
	}

	log.WithField("path", filePath).Info("File processed and saved successfully")
	return nil
}

// Apply adds a processing step to the pipeline.
func (c *Core) Apply(operation string, params map[string]string) error {
	op := &Operation{
		Name:   operation,
		Params: params,
	}
	c.builder.Add(op)
	logger.GetInstance().WithFields(map[string]interface{}{
		"operation": op.Name,
		"params":    op.Params,
	}).Info("Step added to pipeline")
	return nil
}

// SavePipeline saves the current pipeline to a file.
func (c *Core) SavePipeline(filePath string) error {
	log := logger.GetInstance()
	err := c.builder.SaveToFile(filePath)
	if err != nil {
		log.WithField("path", filePath).Errorf("Failed to save pipeline: %v", err)
		return err
	}
	log.WithField("path", filePath).Info("Pipeline saved successfully")
	return nil
}

// LoadPipeline loads a pipeline from a file.
func (c *Core) LoadPipeline(filePath string) error {
	log := logger.GetInstance()
	c.builder.Reset()
	log.Debug("Pipeline builder reset before loading")
	err := c.builder.LoadFromFile(filePath)
	if err != nil {
		log.WithField("path", filePath).Errorf("Failed to load pipeline: %v", err)
		return err
	}
	log.WithField("path", filePath).Info("Pipeline loaded successfully")
	return nil
}

func (c *Core) createStep(operation string, params map[string]string) (func([]byte) ([]byte, error), error) {
	switch operation {
	case "compress":
		compType, err := comp_const.CompressionTypeFromString(strings.ToUpper(params["type"]))
		if err != nil {
			return nil, err
		}
		compressor := compression.NewCompressor(compType)
		return compressor.Compress, nil

	case "decompress":
		compType, err := comp_const.CompressionTypeFromString(strings.ToUpper(params["type"]))
		if err != nil {
			return nil, err
		}
		decompressor := compression.NewDecompressor(compType)
		return decompressor.Decompress, nil

	case "encrypt":
		encType, err := enc_const.EncryptionTypeFromString(strings.ToUpper(params["type"]))
		if err != nil {
			return nil, err
		}
		key, err := ioutil.ReadFile(params["key_file"])
		if err != nil {
			return nil, fmt.Errorf("failed to read key file: %w", err)
		}
		encryptor := encryption.NewEncryptor(encType)
		return func(data []byte) ([]byte, error) {
			return encryptor.Encrypt(data, key)
		}, nil

	case "decrypt":
		encType, err := enc_const.EncryptionTypeFromString(strings.ToUpper(params["type"]))
		if err != nil {
			return nil, err
		}
		key, err := ioutil.ReadFile(params["key_file"])
		if err != nil {
			return nil, fmt.Errorf("failed to read key file: %w", err)
		}
		decryptor := encryption.NewDecryptor(encType)
		return func(data []byte) ([]byte, error) {
			return decryptor.Decrypt(data, key)
		}, nil

	case "calculate":
		calcMethod, err := calc_const.CalculationMethodFromString(strings.ToUpper(params["type"]))
		if err != nil {
			return nil, err
		}
		calculator := calculation.NewCalculator(calcMethod)
		return func(data []byte) ([]byte, error) {
			res, err := calculator.Calculate(string(data))
			return []byte(res), err
		}, nil

	default:
		return nil, fmt.Errorf("unknown operation: %s", operation)
	}
}
