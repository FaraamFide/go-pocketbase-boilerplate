package logger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/FaraamFide/go-pocketbase-boilerplate/backend/internal/config"
)

// InitLogger initializes the global default logger based on the provided log configuration.
// It sets up the log level, format (text or json), and whether to include source code location.
//
// Parameters:
//   - logCfg (config.LogConfig): The logging configuration struct.
//
// Side Effects:
//   - This function sets the global logger for the entire application using slog.SetDefault().
//     After this function is called, any call to slog.Info(), slog.Error(), etc., will use this logger.
func InitLogger(logCfg config.LogConfig) {
	var level slog.Level

	// Convert the log level string from the config to a slog.Level constant.
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
		// If an invalid level is specified in the config, default to Info.
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: logCfg.AddSource, // Useful for debugging: adds file and line number to the log.
	}

	var handler slog.Handler
	// Choose the log handler (and thus format) based on the config.
	switch strings.ToLower(logCfg.Format) {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	case "text":
		fallthrough // If not "json", default to "text".
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	// Set the configured handler as the default for the entire application.
	slog.SetDefault(slog.New(handler))
}
