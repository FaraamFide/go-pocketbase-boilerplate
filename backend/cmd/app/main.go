// Package main is the entry point for the application.
package main

import (
	"flag"
	"log/slog"

	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/config"
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/core"
	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/logger"

	// add if you have new migrations
	// _ "github.com/FaraamFide/go-pocketbase-boilerplate/backend/migrations"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./configs/config.dev.yaml", "path to application configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		slog.Error("Failed to load configuration", "path", configPath, "error", err)
		return
	}

	logger.InitLogger(cfg.Log)
	slog.Info("Configuration and secrets loaded successfully.")

	pbApp, err := core.NewPocketBase(cfg)
	if err != nil {
		slog.Error("Failed to initialize PocketBase", "error", err)
		return
	}
	slog.Info("PocketBase instance created. Starting application...")

	if err := pbApp.Start(); err != nil {
		slog.Error("Application failed to start", "error", err)
	}
}
