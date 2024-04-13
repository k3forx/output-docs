package repository_impl

import (
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryError struct {
	err error
}

func (re RepositoryError) Error() string {
	if e := re.err; e != nil {
		return e.Error()
	}
	return ""
}

func NewRepositoryError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return RepositoryError{
		err: fmt.Errorf("repository error: %w", err),
	}
}
