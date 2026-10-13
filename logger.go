package main

import (
	"log"
	"log/slog"
	"os"
	"strings"
)

func initLogger() *slog.Logger {
	// LOG_LEVEL=debug|info|warn|error (default: info)
	level := slog.LevelInfo

	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		level = slog.LevelDebug
	case "warn", "warning":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	// LOG_FORMAT=console|json (default: console)
	var handler slog.Handler
	switch strings.ToLower(os.Getenv("LOG_FORMAT")) {
	case "json":
		handler = slog.NewJSONHandler(os.Stderr, opts)
	default:
		handler = slog.NewTextHandler(os.Stderr, opts)
	}

	// Set the global slog logger.
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// Route standard log package through slog.
	// slog.SetDefault automatically redirects log.Print* to slog
	// at the level controlled by slog.SetLogLoggerLevel.
	slog.SetLogLoggerLevel(slog.LevelInfo)

	// Remove legacy log prefixes/flags.
	log.SetFlags(0)

	return logger
}
