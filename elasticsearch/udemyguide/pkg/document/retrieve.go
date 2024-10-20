package document

import (
	"context"
	"encoding/json"
	"fmt"

	es "github.com/k3forx/elasticsearchudemy/pkg/client"
	"github.com/k3forx/elasticsearchudemy/pkg/index"
	"github.com/k3forx/elasticsearchudemy/pkg/logger"
)

func RetrieveDocument(ctx context.Context) {
	const docID = "100"

	getResp, err := es.Client.Get(index.ProductIndex, docID).Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed to get doc by id '%s', err: %w", docID, err).Error())
		return
	}
	if !getResp.Found {
		logger.Info(fmt.Sprintf("not found doc by id '%s'", docID))
		return
	}

	var doc ProductDocument
	if err := json.Unmarshal(getResp.Source_, &doc); err != nil {
		logger.Error(fmt.Errorf("failed to unmarshal JSON, err: %w", err).Error())
	}

	logger.Info(fmt.Sprintf("doc: %+v", doc))
}
