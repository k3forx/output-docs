package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"

	"github.com/k3forx/sqlc/tutorial"
)

func run() error {
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:password@/db?parseTime=true")
	if err != nil {
		return err
	}

	queries := tutorial.New(db)

	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	result, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}

	insertedAuthorID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertedAuthorID)

	fetchedAuthor, err := queries.GetAuthor(ctx, -1111)
	if err != nil {
		return err
	}

	log.Println(reflect.DeepEqual(insertedAuthorID, fetchedAuthor.ID))

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
