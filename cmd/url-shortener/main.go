package main

import (
	"os"
	"url-shortener/internal/config"

	"log/slog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config: cleanenv
	config := config.MustLoad()

	// TODO: init logger: slog
	log := setupLogger(config.Env)
	log.Info("starting url-shortener", slog.String("env", config.Env))
	log.Debug("debug message are enabled")

	// TODO: init storage: psql

	// TODO: init router: chi

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {

	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
