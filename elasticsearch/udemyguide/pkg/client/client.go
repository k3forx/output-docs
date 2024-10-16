package es

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

var Client *elasticsearch.TypedClient

func InitElasticsearchClient() error {
	c, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
		APIKey:    "dm8ySFRaSUJ3TEJQR1RoU0dRblA6WEREX1pmVGlUSmFmN1NVRTNiWklpUQ==",
	})
	if err != nil {
		return fmt.Errorf("init client err: %v", err)
	}
	Client = c
	return nil
}
