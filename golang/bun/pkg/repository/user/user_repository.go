package user_repository

import (
	"context"

	user_model "github.com/k3forx/bun/domain/model/user"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (user_model.User, error)
}
