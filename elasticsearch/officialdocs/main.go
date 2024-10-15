package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	_, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		log.Fatalf("failed to create elasticsearch client: %v", err)
	}

}
