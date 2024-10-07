package logger

import (
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}))
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelError,
			}))
	}
	return logger
}
