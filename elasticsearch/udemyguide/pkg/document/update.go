package document

import (
	"context"
	"encoding/json"
	"fmt"

	es "github.com/k3forx/elasticsearchudemy/pkg/client"
	"github.com/k3forx/elasticsearchudemy/pkg/index"
	"github.com/k3forx/elasticsearchudemy/pkg/logger"
)

func UpdateDocument(ctx context.Context) {
	const docID = "100"
	updatedProductDoc := ProductDocument{
		Name: "updated name",
	}

	updateResp, err := es.Client.Update(index.ProductIndex, docID).
		Doc(updatedProductDoc).Do(ctx)
	if err != nil {
		logger.Error(fmt.Errorf("failed to update doc by id '%s', err: %w", docID, err).Error())
		return
	}

	fmt.Println(updateResp.Result)
	var doc ProductDocument
	if err := json.Unmarshal(updateResp.Get.Source_, &doc); err != nil {
		logger.Error(fmt.Errorf("failed to unmarshal JSON, err: %w", err).Error())
		return
	}

	logger.Info(fmt.Sprintf("updated doc: %+v\n", doc))
}
