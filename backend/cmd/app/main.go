package main

import (
	"flag"
	"log/slog"

	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/config"
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/core"
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/logger"
)

// main is the entry point of the application.
// It orchestrates the startup sequence:
// 1. Loads configuration from a YAML file and environment variables.
// 2. Initializes a structured logger based on the loaded configuration.
// 3. Creates and configures a new PocketBase application instance.
// 4. Starts the PocketBase server.
// If any step fails, it logs a critical error and exits.
func main() {

	// 1. Define a command-line flag to specify the configuration file path.
	// The default value is set to the development config, so `make run` works out-of-the-box.
	var configPath string
	flag.StringVar(&configPath, "config", "internal/config/config.dev.yaml", "path to the application configuration file")
	flag.Parse() // Parse the command-line flags

	// 2. Load configuration from the path provided by the flag.
	// The hardcoded path is now gone.
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		slog.Error("Failed to load configuration", "path", configPath, "error", err)
		return
	}

	// 3. Initialize the global logger using the logging configuration from the loaded YAML.
	logger.InitLogger(cfg.Log)
	slog.Info("Configuration and secrets loaded successfully.")

	// 4. Create a new PocketBase application instance with the loaded configuration.
	// This function encapsulates all the setup logic for PocketBase.
	pbApp, err := core.NewPocketBase(cfg)
	if err != nil {
		slog.Error("Failed to initialize PocketBase", "error", err)
		return
	}
	slog.Info("PocketBase instance created. Starting application...")

	// 5. Start the application. This is a blocking call that starts the web server.
	if err := pbApp.Start(); err != nil {
		slog.Error("Application failed to start", "error", err)
	}
}
