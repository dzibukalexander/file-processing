package app

import (
	"github.com/dzibukalexander/file-processing/internal/calculation"
	calc_const "github.com/dzibukalexander/file-processing/internal/calculation/constants"
	compress_const "github.com/dzibukalexander/file-processing/internal/compression/constants"
	"github.com/dzibukalexander/file-processing/internal/compression/gzip"
	"github.com/dzibukalexander/file-processing/internal/encryption/aes"
	encrypt_const "github.com/dzibukalexander/file-processing/internal/encryption/constants"
	"github.com/dzibukalexander/file-processing/internal/procession"
)

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

	// compr := compression.NewCompressor(compressionType)
	// deCompr := compression.NewDecompressor(compressionType)

	// encr := encryption.NewEncryptor(encryptionType)
	// deEncr := encryption.NewDecryptor(encryptionType)

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
