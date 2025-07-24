package logger

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/dzibukalexander/file-processing/internal/config"
	"github.com/sirupsen/logrus"
)

var (
	log  *logrus.Logger
	once sync.Once
)

// GetInstance returns the singleton logger instance.
func GetInstance() *logrus.Logger {
	once.Do(func() {
		log = logrus.New()
		// Default to discarding logs until config is loaded.
		log.SetOutput(ioutil.Discard)
	})
	return log
}

// SetupLogger configures the logger based on the application config.
// This should be called after loading the config.
func SetupLogger() {
	logger := GetInstance()
	if config.AppConfig != nil && config.AppConfig.EnableLogging {
		logger.SetOutput(os.Stdout)
		logger.SetLevel(logrus.InfoLevel)
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	} else {
		logger.SetOutput(ioutil.Discard)
	}
}
