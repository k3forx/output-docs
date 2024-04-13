package user_repository_impl

import (
	"context"

	user_model "github.com/k3forx/bun/domain/model/user"
	user_entity "github.com/k3forx/bun/pkg/db/entity/user"
	user_repository "github.com/k3forx/bun/pkg/repository/user"
	"github.com/k3forx/bun/pkg/repository_impl"
	"github.com/uptrace/bun"
)

var _ user_repository.UserRepository = &userRepositoryImpl{}

type userRepositoryImpl struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *userRepositoryImpl {
	return &userRepositoryImpl{
		db: db,
	}
}

func (impl *userRepositoryImpl) GetByID(ctx context.Context, id int64) (user_model.User, error) {
	user := new(user_entity.User)
	if err := impl.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx); err != nil {
		return user_model.User{}, repository_impl.NewRepositoryError(err)
	}
	return user.Model(), nil
}
