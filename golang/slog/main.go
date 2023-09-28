package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	slog.Info("hello, world")

	// NOTE: msg(第一引数)の後はkey-valueペアで表示される
	slog.Info("hello, world", "user", os.Getenv("USER"))

	// NOTE: ロガーは以下のように明示的に取得できる
	logger := slog.Default()
	logger.Info("hello, world", "user", os.Getenv("USER"))

	// NOTE: TextHandlerを使う
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("hello, world", "user", os.Getenv("USER"))

	// NOTE: JSONHandlerを使う
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello, world", "user", os.Getenv("USER"))

	// NOTE: LogAttrsを使う
	slog.LogAttrs(context.Background(), slog.LevelInfo, "hello, world",
		slog.String("user", os.Getenv("USER")))
}
