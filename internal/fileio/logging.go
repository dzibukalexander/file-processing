package fileio

import (
	"time"

	"github.com/dzibukalexander/file-processing/internal/logger"
)

type loggingFileReader struct {
	reader FileReader
}

func NewLoggingFileReader(reader FileReader) FileReader {
	return &loggingFileReader{reader: reader}
}

func (l *loggingFileReader) Read(filePath string) (data []byte, err error) {
	log := logger.GetInstance().WithField("path", filePath)
	log.Info("Starting file read")

	defer func(begin time.Time) {
		if err != nil {
			log.WithError(err).Error("File read failed")
		} else {
			log.WithFields(map[string]interface{}{
				"duration": time.Since(begin),
				"size":     len(data),
			}).Info("File read finished")
		}
	}(time.Now())

	return l.reader.Read(filePath)
}

type loggingFileWriter struct {
	writer FileWriter
}

func NewLoggingFileWriter(writer FileWriter) FileWriter {
	return &loggingFileWriter{writer: writer}
}

func (l *loggingFileWriter) Write(filePath string, data []byte) (err error) {
	log := logger.GetInstance().WithFields(map[string]interface{}{
		"path": filePath,
		"size": len(data),
	})
	log.Info("Starting file write")

	defer func(begin time.Time) {
		if err != nil {
			log.WithError(err).Error("File write failed")
		} else {
			log.WithField("duration", time.Since(begin)).Info("File write finished")
		}
	}(time.Now())

	return l.writer.Write(filePath, data)
}
