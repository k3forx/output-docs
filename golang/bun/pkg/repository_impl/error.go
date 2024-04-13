package repository_impl

import (
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryError error

func NewRepositoryError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return fmt.Errorf("repository error: %w", err)
}
