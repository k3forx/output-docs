package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

var logger *slog.Logger

func initLogger() {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}
	logger = slog.New(slog.NewTextHandler(os.Stdout, &opts))
}

func main() {
	initLogger()

	ctx := context.Background()

	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
		APIKey:    "dm8ySFRaSUJ3TEJQR1RoU0dRblA6WEREX1pmVGlUSmFmN1NVRTNiWklpUQ==",
	})
	if err != nil {
		logger.Error(fmt.Errorf("init client err: %v", err).Error())
		return
	}

	healthReport := client.HealthReport()
	healthResp, err := healthReport.Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed health report: %v", err).Error())
		return
	}
	logger.Info(fmt.Sprintf("ClusterName: %s", healthResp.ClusterName))
}
