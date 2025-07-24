package config

import (
	"encoding/json"
	"os"
)

// Config holds the application's configuration settings.
type Config struct {
	EnableLogging bool `json:"enable_logging"`
}

// AppConfig is the global configuration instance.
var AppConfig *Config

// LoadConfig loads configuration from a file or uses defaults.
func LoadConfig(path string) error {
	// Default config
	AppConfig = &Config{
		EnableLogging: false,
	}

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			// Config file doesn't exist, use defaults.
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(AppConfig)
	return err
}
