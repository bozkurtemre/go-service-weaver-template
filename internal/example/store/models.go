// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package store

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Example struct {
	ID        int64
	Message   string
	CreatedAt pgtype.Timestamptz
}
