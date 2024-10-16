package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func InitLogger() {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &opts))
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}
