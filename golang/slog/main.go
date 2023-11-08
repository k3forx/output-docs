package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	var a any
	a = "pana"
	fmt.Println(a)
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(h)
	logger.Info("some message", "userId", 123, "name", "pana")
	slog.Int()

	// logger = logger.With("userId", 123, "name", "pana")
	// logger.Info("some message")
	// logger.WithGroup("grouped").Info("some message", "userId", 123, "name", "pana")
}
