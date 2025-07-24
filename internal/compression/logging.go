package compression

import (
	"time"

	"github.com/dzibukalexander/file-processing/internal/logger"
)

type loggingCompressor struct {
	compressor Compressor
}

func NewLoggingCompressor(compressor Compressor) Compressor {
	return &loggingCompressor{compressor: compressor}
}

func (l *loggingCompressor) Compress(data []byte) (result []byte, err error) {
	log := logger.GetInstance().WithField("input_size", len(data))
	log.Info("Starting compression")

	defer func(begin time.Time) {
		if err != nil {
			log.WithError(err).Error("Compression failed")
		} else {
			log.WithFields(map[string]interface{}{
				"duration":    time.Since(begin),
				"output_size": len(result),
				"ratio":       float64(len(result)) / float64(len(data)),
			}).Info("Compression finished")
		}
	}(time.Now())

	return l.compressor.Compress(data)
}

type loggingDecompressor struct {
	decompressor Decompressor
}

func NewLoggingDecompressor(decompressor Decompressor) Decompressor {
	return &loggingDecompressor{decompressor: decompressor}
}

func (l *loggingDecompressor) Decompress(data []byte) (result []byte, err error) {
	log := logger.GetInstance().WithField("input_size", len(data))
	log.Info("Starting decompression")

	defer func(begin time.Time) {
		if err != nil {
			log.WithError(err).Error("Decompression failed")
		} else {
			log.WithFields(map[string]interface{}{
				"duration":    time.Since(begin),
				"output_size": len(result),
			}).Info("Decompression finished")
		}
	}(time.Now())

	return l.decompressor.Decompress(data)
}
