// Package logger provides a helper for initializing the global structured logger.
package logger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/config"
)

// InitLogger configures and sets the global default logger for the application
// based on the provided configuration.
func InitLogger(logCfg config.LogConfig) {
	var level slog.Level

	switch strings.ToLower(logCfg.Level) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: logCfg.AddSource,
	}

	var handler slog.Handler

	switch strings.ToLower(logCfg.Format) {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	// Set the configured handler as the default for the entire application.
	slog.SetDefault(slog.New(handler))
}
