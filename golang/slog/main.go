package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	// slog.Info("hello, world")

	// NOTE: msg(第一引数)の後はkey-valueペアで表示される
	// slog.Info("hello, world", "user", os.Getenv("USER"))

	// NOTE: ロガーは以下のように明示的に取得できる
	// logger := slog.Default()
	// logger.Info("hello, world", "user", os.Getenv("USER"))

	// NOTE: TextHandlerを使う
	// logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	// logger.Info("hello, world", "user", os.Getenv("USER"))

	// NOTE: JSONHandlerを使う
	opts := slog.HandlerOptions{
		AddSource: true,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &opts))
	logger.Info("hello\nworld", "user", os.Getenv("USER"))

	// NOTE: LogAttrsを使う
	// slog.LogAttrs(context.Background(), slog.LevelInfo, "hello, world",
	// 	slog.String("user", os.Getenv("USER")))

	slog.SetDefault(logger)
	log.Println("log from log pkg")

	var progLevel = new(slog.LevelVar)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: progLevel})
	logger = slog.New(h)
	slog.SetDefault(logger)
	progLevel.Set(slog.LevelError)
	logger.Info("info")

	logger.With()
}
