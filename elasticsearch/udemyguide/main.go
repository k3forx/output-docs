package main

import (
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	es "github.com/k3forx/elasticsearchudemy/pkg/client"
	"github.com/k3forx/elasticsearchudemy/pkg/document"
	"github.com/k3forx/elasticsearchudemy/pkg/index"
	"github.com/k3forx/elasticsearchudemy/pkg/logger"
)

func main() {
	logger.InitLogger()

	if err := es.InitElasticsearchClient(); err != nil {
		logger.Error(err.Error())
		return
	}

	ctx := context.Background()

	healthReport := es.Client.HealthReport()
	healthResp, err := healthReport.Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed health report: %v", err).Error())
		return
	}
	logger.Info(fmt.Sprintf("ClusterName: %s", healthResp.ClusterName))

	// --------------------------------------------------

	const pagesIndex = "pages"
	createResp, err := index.GenCreateIndexRequest(pagesIndex).Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed to create index: %s, err: %w", pagesIndex, err).Error())
		return
	}
	if createResp.Acknowledged {
		logger.Info(fmt.Sprintf("success to create index: %s", pagesIndex))
	}

	deleteResp, err := es.Client.Indices.Delete(pagesIndex).Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed to delete index: %s, err: %w", pagesIndex, err).Error())
		return
	}
	if deleteResp.Acknowledged {
		logger.Info(fmt.Sprintf("success to delete index: %s", pagesIndex))
	}

	// --------------------------------------------------

	_, err = index.GenDeleteIndexRequest(index.ProductIndex).Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed to delete index: %s, err: %w", pagesIndex, err).Error())
		return
	}
	if deleteResp.Acknowledged {
		logger.Info(fmt.Sprintf("success to delete index: %s", index.ProductIndex))
	}

	createResp, err = index.GenCreateIndexRequest(index.ProductIndex).Settings(
		&types.IndexSettings{
			NumberOfShards: "2",
		},
	).Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed to create index: %s, err: %w", index.ProductIndex, err).Error())
		return
	}
	if createResp.Acknowledged {
		logger.Info(fmt.Sprintf("success to create index: %s", index.ProductIndex))
	}

	document.IndexDocument(ctx)
}
