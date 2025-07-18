package app

import (
	"flag"
	"fmt"
	"strings"
)

type cliParams struct {
	inputFile      string
	outputFile     string
	calcMethodStr  string
	compressionStr string
	encryptionStr  string
}

func parseFlags() (*cliParams, error) {
	inputFile := flag.String("input", "", "Input file path (supports .txt, .json, .xml, .yaml, .html)")
	outputFile := flag.String("output", "", "Output file path (supports .txt, .json, .xml, .yaml, .html)")
	calcMethodStr := flag.String("calc", "REGEX", "Calculation method (REGEX, PARSER, LIBRARY)")
	compressionStr := flag.String("compression", "NONE", "Compression type (NONE, ZIP, GZIP)")
	encryptionStr := flag.String("encryption", "NONE", "Encryption type (NONE, AES, RSA)")

	flag.Parse()

	*calcMethodStr = strings.ToUpper(*calcMethodStr)
	*compressionStr = strings.ToUpper(*compressionStr)
	*encryptionStr = strings.ToUpper(*encryptionStr)

	return &cliParams{
		inputFile:      *inputFile,
		outputFile:     *outputFile,
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
