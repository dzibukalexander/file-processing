package procession

import (
	"fmt"

	"github.com/dzibukalexander/file-processing/internal/calculation"
	"github.com/dzibukalexander/file-processing/internal/compression"
	"github.com/dzibukalexander/file-processing/internal/encryption"
	"github.com/dzibukalexander/file-processing/internal/fileio"
)

type FileProcessor struct {
	calculator   calculation.Calculator
	compressor   compression.Compressor
	decompressor compression.Decompressor
	encryptor    encryption.Encryptor
	decryptor    encryption.Decryptor
}

func NewFileProcessor(
	calc calculation.Calculator,
	comp compression.Compressor,
	decomp compression.Decompressor,
	enc encryption.Encryptor,
	dec encryption.Decryptor,
) *FileProcessor {
	return &FileProcessor{
		calculator:   calc,
		compressor:   comp,
		decompressor: decomp,
		encryptor:    enc,
		decryptor:    dec,
	}
}

type Operation func([]byte) ([]byte, error)

func (p *FileProcessor) BuildPipeline(operations ...Operation) Operation {
	return func(data []byte) ([]byte, error) {
		var err error
		for _, op := range operations {
			data, err = op(data)
			if err != nil {
				return nil, err
			}
		}
		return data, nil
	}
}

func (p *FileProcessor) ReadOperation() Operation {
	return func(data []byte) ([]byte, error) {
		return data, nil // Просто пассим данные, чтение делается отдельно
	}
}

func (p *FileProcessor) DecompressOperation() Operation {
	return func(data []byte) ([]byte, error) {
		if p.decompressor != nil {
			return p.decompressor.Decompress(data)
		}
		return data, nil
	}
}

func (p *FileProcessor) DecryptOperation() Operation {
	return func(data []byte) ([]byte, error) {
		if p.decryptor != nil {
			return p.decryptor.Decrypt(data, getEncryptionKey())
		}
		return data, nil
	}
}

// для алгоритмов шифрования/де нужен ключ
func getEncryptionKey() []byte {
	return []byte("secure-encryption-key-32-bytes-long")
}

func (p *FileProcessor) CalculateOperation() Operation {
	return func(data []byte) ([]byte, error) {
		result, err := p.calculator.Calculate(string(data))
		return []byte(result), err
	}
}

func (p *FileProcessor) EncryptOperation() Operation {
	return func(data []byte) ([]byte, error) {
		if p.encryptor != nil {
			return p.encryptor.Encrypt(data, getEncryptionKey())
		}
		return data, nil
	}
}

func (p *FileProcessor) CompressOperation() Operation {
	return func(data []byte) ([]byte, error) {
		if p.compressor != nil {
			return p.compressor.Compress(data)
		}
		return data, nil
	}
}

func (p *FileProcessor) WriteOperation() Operation {
	return func(data []byte) ([]byte, error) {
		return data, nil // Запись делается отдельно
	}
}

func (p *FileProcessor) ProcessFile(
	reader fileio.FileReader,
	writer fileio.FileWriter,
	inputPath, outputPath string,
	pipeline Operation,
) error {
	content, err := reader.Read(inputPath)
	if err != nil {
		return fmt.Errorf("reading error: %v", err)
	}

	processed, err := pipeline(content)
	if err != nil {
		return fmt.Errorf("processing error: %v", err)
	}

	err = writer.Write(outputPath, processed)
	if err != nil {
		return fmt.Errorf("writing error: %v", err)
	}

	return nil
}
