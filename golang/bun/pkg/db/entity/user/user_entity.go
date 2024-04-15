package user_entity

import (
	user_model "github.com/k3forx/bun/domain/model/user"
)

type User struct {
	ID    int64  `bun:"id,pk,autoincrement"`
	Name  string `bun:"name"`
	Email string `bun:"email"`
	// CreatedAt time.Time `bun:"created_at"` // usersテーブルにはないフィールドを追加してもエラーにならない
}

func (e User) Model() user_model.User {
	return user_model.User{
		ID:   e.ID,
		Name: e.Name,
	}
}
