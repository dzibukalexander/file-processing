package app

import (
	"fmt"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() error {
	params, err := parseFlags()
	if err != nil {
		return fmt.Errorf("error parsing flags: %w", err)
	}

	if err := validateInputOutput(params.inputFile, params.outputFile); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	inputType, outputType, err := getFileTypes(params.inputFile, params.outputFile)
	if err != nil {
		return fmt.Errorf("error determining file types: %w", err)
	}

	calcMethod, err := parseCalculationMethod(params.calcMethodStr)
	if err != nil {
		return fmt.Errorf("error parsing calculation method: %w", err)
	}

	compressionType, err := parseCompressionType(params.compressionStr)
	if err != nil {
		return fmt.Errorf("error parsing compression type: %w", err)
	}

	encryptionType, err := parseEncryptionType(params.encryptionStr)
	if err != nil {
		return fmt.Errorf("error parsing encryption type: %w", err)
	}

	processor, err := setupFileProcessor(calcMethod, compressionType, encryptionType)
	if err != nil {
		return fmt.Errorf("error setting up processor: %w", err)
	}

	if err := processFile(processor, inputType, outputType, params.inputFile, params.outputFile); err != nil {
		return fmt.Errorf("error processing file: %w", err)
	}

	return nil
}
