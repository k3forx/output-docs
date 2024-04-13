package user_repository_impl_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/uptrace/bun"

	user_model "github.com/k3forx/bun/domain/model/user"
	user_entity "github.com/k3forx/bun/pkg/db/entity/user"
	user_repository "github.com/k3forx/bun/pkg/repository/user"
	user_repository_impl "github.com/k3forx/bun/pkg/repository_impl/user"
	"github.com/k3forx/bun/pkg/test_util"
)

var (
	db   *bun.DB
	impl user_repository.UserRepository
)

func TestMain(m *testing.M) {
	db = test_util.SetupTestDB(m)
	impl = user_repository_impl.NewUserRepository(db)
	os.Exit(m.Run())
}

func TestUserRepositoryImpl_GetByID(t *testing.T) {
	ctx := context.Background()
	test_util.LoadFixtures(
		t, ctx, db,
		"../testdata/7FD21D5D-CEC9-4C9D-9617-DA15CE423E28.yaml",
		(*user_entity.User)(nil),
	)

	cases := map[string]struct {
		id       int64
		expected user_model.User
	}{
		"id=1の場合はID=1のユーザーを取得する": {
			id: 1,
			expected: user_model.User{
				ID: 1,
			},
		},
		"存在しないIDなので空の構造体が返ってくる": {
			id:       0,
			expected: user_model.User{},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := impl.GetByID(ctx, c.id)
			if err != nil {
				t.Errorf("err should be nil, but got: %v", err)
				return
			}
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
