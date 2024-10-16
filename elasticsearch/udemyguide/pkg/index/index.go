package index

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/delete"
	es "github.com/k3forx/elasticsearchudemy/pkg/client"
)

const (
	ProductIndex = "products"
)

func GenCreateIndexRequest(indexName string) *create.Create {
	return es.Client.Indices.Create(indexName)
}

func GenDeleteIndexRequest(indexName string) *delete.Delete {
	return es.Client.Indices.Delete(indexName)
}
