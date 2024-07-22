package main

import (
	"AnkiConverter/internal/config"
	"AnkiConverter/internal/dictionary"
	libretranslate "AnkiConverter/internal/libretranslate"
	"encoding/json"
	"log/slog"
	"os"
	"strings"
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

	translated, err := libretranslate.LibreTranslate("Человек",
		"ru",
		"en",
		"text",
		3,
		"")

	if err != nil {
		log.Error("failed to translate text %s", err.Error())
	}

	text, _ := json.Marshal(translated)
	log.Info(string(text))

	translated.TranslatedText = strings.Trim(translated.TranslatedText, "!@#$%^&*()_-=+`~")
	texts := strings.Split(translated.TranslatedText, " ")
	for _, word := range texts {
		respDict, err := dictionary.GetDictionary(word)
		if err != nil {
			log.Error("failed to get dictionary %s", err.Error())
		}

		res, _ := json.Marshal(respDict)
		log.Info(string(res))
	}

}
