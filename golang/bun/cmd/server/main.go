package main

import (
	"context"
	"fmt"
	"log"

	"github.com/k3forx/bun/pkg/db"
)

func main() {
	ctx := context.Background()

	db, err := db.Setup()
	if err != nil {
		panic(err)
	}

	var num int
	if err := db.QueryRowContext(ctx, "SELECT 1").Scan(&num); err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
}
