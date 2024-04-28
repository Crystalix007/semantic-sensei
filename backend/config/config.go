// Package config provides the configuration specification.
package config

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

// DefaultConfigPath is the default path to the configuration file.
const DefaultConfigPath = "/app/config.yaml"

// ConfigEnvVar is the environment variable that can be used to set the path to
// the configuration file.
const ConfigEnvVar = "SEMANTIC_SENSEI_CONFIG_PATH"

// Config represents the configuration settings for the application.
type Config struct {
	Database Database `yaml:"database"`
}

// New creates a new instance of Config by reading the configuration file.
// It first checks if the config path environment variable is set, and if so,
// uses that as the path to the configuration file.
// If the environment variable is not set, it uses the default path.
// Then the config file is read, and unmarshalled.
// Returns a pointer to the Config struct and any error encountered during the
// process.
func New() (*Config, error) {
	cfgFilePath := DefaultConfigPath

	if envPath := os.Getenv(ConfigEnvVar); envPath != "" {
		slog.Debug("config: using config file from environment variable", slog.String("path", envPath))

		cfgFilePath = envPath
	}

	cfgFile, err := os.ReadFile(cfgFilePath)
	if err != nil {
		return nil, fmt.Errorf(
			"config: error reading config file: %w",
			err,
		)
	}

	var cfg Config

	if err := yaml.Unmarshal(cfgFile, &cfg); err != nil {
		return nil, fmt.Errorf(
			"config: error unmarshalling config file: %w",
			err,
		)
	}

	return &cfg, nil
}
