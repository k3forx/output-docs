package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/delete"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

var (
	logger *slog.Logger
	client *elasticsearch.TypedClient
)

func initLogger() {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &opts))
}

func initElasticsearchClient() error {
	c, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
		APIKey:    "dm8ySFRaSUJ3TEJQR1RoU0dRblA6WEREX1pmVGlUSmFmN1NVRTNiWklpUQ==",
	})
	if err != nil {
		return fmt.Errorf("init client err: %v", err)
	}
	client = c
	return nil
}

func main() {
	initLogger()

	if err := initElasticsearchClient(); err != nil {
		logger.Error(err.Error())
		return
	}

	ctx := context.Background()

	healthReport := client.HealthReport()
	healthResp, err := healthReport.Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed health report: %v", err).Error())
		return
	}
	logger.Info(fmt.Sprintf("ClusterName: %s", healthResp.ClusterName))

	// --------------------------------------------------

	const pagesIndex = "pages"
	createResp, err := genCreateIndexRequest(pagesIndex).Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed to create index: %s, err: %w", pagesIndex, err).Error())
		return
	}
	if createResp.Acknowledged {
		logger.Info(fmt.Sprintf("success to create index: %s", pagesIndex))
	}

	deleteResp, err := client.Indices.Delete(pagesIndex).Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed to delete index: %s, err: %w", pagesIndex, err).Error())
		return
	}
	if deleteResp.Acknowledged {
		logger.Info(fmt.Sprintf("success to delete index: %s", pagesIndex))
	}

	// --------------------------------------------------

	const productIndex = "products"
	_, err = genDeleteIndexRequest(productIndex).Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed to delete index: %s, err: %w", pagesIndex, err).Error())
		return
	}
	if deleteResp.Acknowledged {
		logger.Info(fmt.Sprintf("success to delete index: %s", productIndex))
	}

	createResp, err = genCreateIndexRequest(productIndex).Settings(
		&types.IndexSettings{
			NumberOfShards: "2",
		},
	).Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed to create index: %s, err: %w", productIndex, err).Error())
		return
	}
	if createResp.Acknowledged {
		logger.Info(fmt.Sprintf("success to create index: %s", productIndex))
	}
}

func genCreateIndexRequest(indexName string) *create.Create {
	return client.Indices.Create(indexName)
}

func genDeleteIndexRequest(indexName string) *delete.Delete {
	return client.Indices.Delete(indexName)
}
