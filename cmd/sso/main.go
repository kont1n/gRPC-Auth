package main

import (
	"log/slog"

	"gRPC-Auth/internal/config"
	"gRPC-Auth/internal/utils/logger"
)

var err error

func main() {

	cfg := config.Load()

	log := logger.SetupLogger(cfg.Env)

	log.Info("Starting application",
		slog.String("env", cfg.Env),
	)

	// TODO: Инициализировать приложение
}
