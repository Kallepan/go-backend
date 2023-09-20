package config

import (
	"log/slog"
	"os"
)

func InitlLog() {
	opts := &slog.HandlerOptions{
		Level:     getLoggerLevel(),
		AddSource: true,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}

func getLoggerLevel() slog.Level {
	value := os.Getenv("LOG_LEVEL")

	switch value {
	case "DEBUG":
		return slog.LevelDebug
	case "TRACE":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
