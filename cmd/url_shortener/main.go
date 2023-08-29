package main

import (
	"log/slog"
	"os"
	"url_shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	log := setupLog(cfg.Env)

	log.Info("starting service", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")
}

func setupLog(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case config.EnvLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.EnvDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.EnvProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
