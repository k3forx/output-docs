package document

import (
	"context"
	"fmt"

	es "github.com/k3forx/elasticsearchudemy/pkg/client"
	"github.com/k3forx/elasticsearchudemy/pkg/index"
	"github.com/k3forx/elasticsearchudemy/pkg/logger"
)

type ProductDocument struct {
	Name    string `json:"name"`
	Price   int    `json:"price"`
	InStock int    `json:"in_stock"`
}

func IndexDocument(ctx context.Context) {
	doc := ProductDocument{
		Name:    "Coffee Maker",
		Price:   64,
		InStock: 10,
	}
	createResp, err := es.Client.Index(index.ProductIndex).Request(doc).Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed to create document. err: %w", err).Error())
		return
	}

	logger.Info(fmt.Sprintf("id: %s", createResp.Id_))

	doc = ProductDocument{
		Name:    "Toaster",
		Price:   49,
		InStock: 4,
	}
	createResp, err = es.Client.Index(index.ProductIndex).
		Id("100").Request(doc).Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed to create document. err: %w", err).Error())
		return
	}

	logger.Info(fmt.Sprintf("specified id: %s", createResp.Id_))
}
