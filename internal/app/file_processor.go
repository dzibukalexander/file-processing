package app

import (
	"fmt"

	"github.com/dzibukalexander/file-processing/internal/fileio"
	file_const "github.com/dzibukalexander/file-processing/internal/fileio/constants"
	"github.com/dzibukalexander/file-processing/internal/procession"
)

func getFileTypes(inputFile, outputFile string) (inputType, outputType file_const.FileType, err error) {
	inputType, err = file_const.FileTypeFromExtension(inputFile)
	if err != nil {
		return "", "", fmt.Errorf("invalid input file type: %w", err)
	}

	outputType, err = file_const.FileTypeFromExtension(outputFile)
	if err != nil {
		return "", "", fmt.Errorf("invalid output file type: %w", err)
	}

	return inputType, outputType, nil
}

func processFile(
	processor *procession.FileProcessor,
	inputType file_const.FileType,
	outputType file_const.FileType,
	inputPath string,
	outputPath string,
) error {
	reader := fileio.NewFileReader(inputType)
	writer := fileio.NewWriter(outputType)

	pipeline := processor.BuildPipeline(
		processor.CalculateOperation(),
	)

	return processor.ProcessFile(reader, writer, inputPath, outputPath, pipeline)
}
