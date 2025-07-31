// Package config provides structures and functions for loading application configuration
// from YAML files and environment variables.
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// LogConfig holds the configuration for the application logger.
type LogConfig struct {
	Level     string `yaml:"level"`
	Format    string `yaml:"format"`
	AddSource bool   `yaml:"addSource"`
}

// ServerConfig holds the configuration for the server.
type ServerConfig struct {
	PocketbaseHost string `yaml:"pocketbaseHost"`
	PocketbasePort string `yaml:"pocketbasePort"`
	AppUrl         string `yaml:"appUrl"`
	DataDir        string `yaml:"dataDir"`
}

// SecretsConfig holds application secrets loaded from the environment.
type SecretsConfig struct {
	EncryptionKey string
}

// Config holds all configuration parameters for the application.
type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Log     LogConfig     `yaml:"log"`
	Secrets SecretsConfig // Loaded from environment, not YAML.
}

func LoadConfig(configPath string) (*Config, error) {

	_ = godotenv.Load()

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("could not read config file %s: %w", configPath, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("could not unmarshal yaml: %w", err)
	}

	cfg.Secrets.EncryptionKey = os.Getenv("POCKETBASE_ENCRYPTION_KEY")
	if cfg.Secrets.EncryptionKey == "" {
		fmt.Println("WARNING: POCKETBASE_ENCRYPTION_KEY environment variable not set.")
	}

	return &cfg, nil
}
