package main

import (
	"context"
	"fmt"
	"log"

	"github.com/k3forx/bun/pkg/db"
	user_repository_impl "github.com/k3forx/bun/pkg/repository_impl/user"
)

func main() {
	ctx := context.Background()

	db, err := db.Setup()
	if err != nil {
		panic(err)
	}

	userRepo := user_repository_impl.NewUserRepository(db)
	user, err := userRepo.GetByID(ctx, 1)
	if err != nil {
		log.Printf("err: %+v\n", err)
		return
	}
	fmt.Printf("user: %+v\n", user)
}
