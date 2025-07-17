package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dzibukalexander/file-processing/internal/calculation"
	calc_const "github.com/dzibukalexander/file-processing/internal/calculation/constants"
	compress_const "github.com/dzibukalexander/file-processing/internal/compression/constants"
	"github.com/dzibukalexander/file-processing/internal/compression/gzip"
	"github.com/dzibukalexander/file-processing/internal/encryption/aes"
	encrypt_const "github.com/dzibukalexander/file-processing/internal/encryption/constants"
	"github.com/dzibukalexander/file-processing/internal/fileio"
	file_const "github.com/dzibukalexander/file-processing/internal/fileio/constants"
	"github.com/dzibukalexander/file-processing/internal/procession"
)

func main() {
	params, err := parseFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := validateInputOutput(params.inputFile, params.outputFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileTypes, err := parseFileTypes(params.inputTypeStr, params.outputTypeStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	calcMethod, err := parseCalculationMethod(params.calcMethodStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	compressionType, err := parseCompressionType(params.compressionStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	encryptionType, err := parseEncryptionType(params.encryptionStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	processor, err := setupFileProcessor(calcMethod, compressionType, encryptionType)
	if err != nil {
		fmt.Printf("Error setting up processor: %v\n", err)
		os.Exit(1)
	}

	if err := processFile(processor, fileTypes.inputType, fileTypes.outputType, params.inputFile, params.outputFile); err != nil {
		fmt.Printf("Error processing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("File processed successfully")
}

type cliParams struct {
	inputFile      string
	outputFile     string
	inputTypeStr   string
	outputTypeStr  string
	calcMethodStr  string
	compressionStr string
	encryptionStr  string
}

func parseFlags() (*cliParams, error) {
	inputFile := flag.String("input", "", "Input file path")
	outputFile := flag.String("output", "", "Output file path")
	inputTypeStr := flag.String("input-type", "TEXT", "Input file type (TEXT, JSON, XML, YAML, HTML)")
	outputTypeStr := flag.String("output-type", "TEXT", "Output file type (TEXT, JSON, XML, YAML, HTML)")
	calcMethodStr := flag.String("calc", "REGEX", "Calculation method (REGEX, PARSER, LIBRARY)")
	compressionStr := flag.String("compression", "NONE", "Compression type (NONE, ZIP, GZIP)")
	encryptionStr := flag.String("encryption", "NONE", "Encryption type (NONE, AES, RSA)")

	flag.Parse()

	*inputTypeStr = strings.ToUpper(*inputTypeStr)
	*outputTypeStr = strings.ToUpper(*outputTypeStr)
	*calcMethodStr = strings.ToUpper(*calcMethodStr)
	*compressionStr = strings.ToUpper(*compressionStr)
	*encryptionStr = strings.ToUpper(*encryptionStr)

	return &cliParams{
		inputFile:      *inputFile,
		outputFile:     *outputFile,
		inputTypeStr:   *inputTypeStr,
		outputTypeStr:  *outputTypeStr,
		calcMethodStr:  *calcMethodStr,
		compressionStr: *compressionStr,
		encryptionStr:  *encryptionStr,
	}, nil
}

func validateInputOutput(inputFile, outputFile string) error {
	if inputFile == "" || outputFile == "" {
		return fmt.Errorf("input and output file paths are required")
	}
	return nil
}

type fileTypes struct {
	inputType  file_const.FileType
	outputType file_const.FileType
}

func parseFileTypes(inputTypeStr, outputTypeStr string) (*fileTypes, error) {
	inputType, err := file_const.FileTypeFromString(inputTypeStr)
	if err != nil {
		return nil, fmt.Errorf("invalid input type: %v", err)
	}

	outputType, err := file_const.FileTypeFromString(outputTypeStr)
	if err != nil {
		return nil, fmt.Errorf("invalid output type: %v", err)
	}

	return &fileTypes{
		inputType:  inputType,
		outputType: outputType,
	}, nil
}

func parseCalculationMethod(calcMethodStr string) (calc_const.CalculationMethod, error) {
	return calc_const.CalculationMethodFromString(calcMethodStr)
}

func parseCompressionType(compressionStr string) (compress_const.CompressionType, error) {
	return compress_const.CompressionTypeFromString(compressionStr)
}

func parseEncryptionType(encryptionStr string) (encrypt_const.EncryptionType, error) {
	return encrypt_const.EncryptionTypeFromString(encryptionStr)
}

func setupFileProcessor(
	calcMethod calc_const.CalculationMethod,
	compressionType compress_const.CompressionType,
	encryptionType encrypt_const.EncryptionType,
) (*procession.FileProcessor, error) {
	calc := calculation.NewCalculator(calcMethod)

	// To do
	gzipCompressor := &gzip.GzipCompressor{}
	gzipDecompressor := &gzip.GzipDecompressor{}
	aesEncryptor := &aes.AESEncryptor{}
	aesDecryptor := &aes.AESDecryptor{}

	return procession.NewFileProcessor(
		calc,
		gzipCompressor,
		gzipDecompressor,
		aesEncryptor,
		aesDecryptor,
	), nil
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
		//processor.DecompressOperation(),
		// processor.DecryptOperation(),
		processor.CalculateOperation(),
		// processor.EncryptOperation(),
		// processor.CompressOperation(),
	)

	return processor.ProcessFile(reader, writer, inputPath, outputPath, pipeline)
}
