package log

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func InitLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}

	handler := tint.NewHandler(os.Stderr, &tint.Options{
		AddSource:  true,
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	})

	if os.Getenv("APP_ENV") == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)

	return logger
}

func InitDefaultLogger() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(handler)

	return logger
}
