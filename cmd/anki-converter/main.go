package main

import (
	"AnkiConverter/internal/config"
	"AnkiConverter/internal/libre-translate"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info(
		"starting AssemblingAnki",
		slog.String("env", cfg.Env),
		slog.String("version", "0.0.1a"),
	)
	log.Debug("debug messages are enabled")

	translated, err := libre_translate.LibreTranslate("Я молодой и красивый заяц",
		"ru",
		"en",
		"text",
		3,
		"")

	if err != nil {
		// TODO: logging
		fmt.Println(err)
	}

	res, _ := json.Marshal(translated)

	println(string(res))
}
