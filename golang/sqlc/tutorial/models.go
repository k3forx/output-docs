// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package tutorial

import (
	"database/sql"
	"time"
)

type Author struct {
	ID        int64
	Name      string
	Bio       sql.NullString
	CreatedAt time.Time
}