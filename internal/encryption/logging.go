package encryption

import (
	"time"

	"github.com/dzibukalexander/file-processing/internal/logger"
)

type loggingEncryptor struct {
	encryptor Encryptor
}

func NewLoggingEncryptor(encryptor Encryptor) Encryptor {
	return &loggingEncryptor{encryptor: encryptor}
}

func (l *loggingEncryptor) Encrypt(data []byte, key []byte) (result []byte, err error) {
	log := logger.GetInstance().WithField("input_size", len(data))
	log.Info("Starting encryption")

	defer func(begin time.Time) {
		if err != nil {
			log.WithError(err).Error("Encryption failed")
		} else {
			log.WithFields(map[string]interface{}{
				"duration":    time.Since(begin),
				"output_size": len(result),
			}).Info("Encryption finished")
		}
	}(time.Now())

	return l.encryptor.Encrypt(data, key)
}

type loggingDecryptor struct {
	decryptor Decryptor
}

func NewLoggingDecryptor(decryptor Decryptor) Decryptor {
	return &loggingDecryptor{decryptor: decryptor}
}

func (l *loggingDecryptor) Decrypt(data []byte, key []byte) (result []byte, err error) {
	log := logger.GetInstance().WithField("input_size", len(data))
	log.Info("Starting decryption")

	defer func(begin time.Time) {
		if err != nil {
			log.WithError(err).Error("Decryption failed")
		} else {
			log.WithFields(map[string]interface{}{
				"duration":    time.Since(begin),
				"output_size": len(result),
			}).Info("Decryption finished")
		}
	}(time.Now())

	return l.decryptor.Decrypt(data, key)
}
