package repository_impl

import (
	"database/sql"
	"errors"
	"fmt"
)

type repositoryError struct {
	err error
}

func (re repositoryError) Error() string {
	if e := re.err; e != nil {
		return e.Error()
	}
	return ""
}

type RepositoryErrorUnexpectedError struct {
	repositoryError
}

type RepositoryErrorNoRowsError struct {
	repositoryError
}

type optConfig struct {
	notIgnoreErrNoRows bool
}

type RepositoryErrorOption interface {
	apply(optConfig) optConfig
}

type RepositoryErrorOptionNotIgnoreErrNoRows struct{}

func (opt RepositoryErrorOptionNotIgnoreErrNoRows) apply(cfg optConfig) optConfig {
	cfg.notIgnoreErrNoRows = true
	return cfg
}

func NewRepositoryError(err error, opts ...RepositoryErrorOption) error {
	baseErr := repositoryError{err: fmt.Errorf("レポジトリエラー: %w", err)}

	var cfg optConfig
	for _, opt := range opts {
		cfg = opt.apply(cfg)
	}

	if errors.Is(err, sql.ErrNoRows) {
		if cfg.notIgnoreErrNoRows {
			return RepositoryErrorNoRowsError{repositoryError: baseErr}
		}
		return nil
	}

	if err != nil {
		return RepositoryErrorUnexpectedError{baseErr}
	}
	return nil
}
