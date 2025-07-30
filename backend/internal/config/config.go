package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// LogConfig holds the configuration for the application logger.
type LogConfig struct {
	Level     string `yaml:"level"`     // Logging level (e.g., "debug", "info", "warn", "error").
	Format    string `yaml:"format"`    // Log output format ("text" or "json").
	AddSource bool   `yaml:"addSource"` // Whether to include the source file and line number in logs.
}

// ServerConfig holds the configuration for the server.
type ServerConfig struct {
	PocketbaseHost string `yaml:"pocketbaseHost"` // Host for the PocketBase server (e.g., "127.0.0.1").
	PocketbasePort string `yaml:"pocketbasePort"` // Port for the PocketBase server (e.g., "8090").
	AppUrl         string `yaml:"appUrl"`         // Frontend URL, used for links in emails, etc.
	DataDir        string `yaml:"dataDir"`        // Directory to store PocketBase data (e.g., "pb_data").
}

// SecretsConfig holds application secrets, typically loaded from environment variables.
type SecretsConfig struct {
	// == IMPORTANT ==
	// This key is used for encrypting and decrypting auth tokens and other sensitive data.
	// It MUST be a 32-character string.
	EncryptionKey string
}

// Config holds all configuration parameters for the application.
type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Log     LogConfig     `yaml:"log"`
	Secrets SecretsConfig // This part is not loaded from YAML, but from the environment.
}

// LoadConfig loads configuration from a specified YAML file and supplements it
// with secrets from environment variables (via a .env file or the system environment).
// It does not use a logger, as the logger itself has not been initialized at this stage.
func LoadConfig(configPath string) (*Config, error) {
	// --- 1. Load secrets from .env file into the environment ---
	// We load this first so that environment variables are available for the next steps.
	// It's safe to ignore the error if the .env file doesn't exist, as variables
	// might be set directly in the environment (e.g., in a Docker container or CI/CD pipeline).
	_ = godotenv.Load()

	// --- 2. Read and unmarshal the main configuration file ---
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("could not read config file %s: %w", configPath, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("could not unmarshal yaml: %w", err)
	}

	// --- 3. Load secrets from the environment into the Config struct ---
	cfg.Secrets.EncryptionKey = os.Getenv("POCKETBASE_ENCRYPTION_KEY")
	if cfg.Secrets.EncryptionKey == "" {
		// NOTE: This is a non-fatal warning. PocketBase will start but might complain
		// or generate a default key if this is missing. It's better to be explicit.
		fmt.Println("WARNING: POCKETBASE_ENCRYPTION_KEY environment variable not set.")
	}

	return &cfg, nil
}
