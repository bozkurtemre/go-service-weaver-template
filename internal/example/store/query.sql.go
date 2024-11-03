// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createExample = `-- name: CreateExample :one
INSERT INTO examples (
  id, message, created_at
) VALUES (
  $1, $2, $3
) RETURNING id, message, created_at
`

type CreateExampleParams struct {
	ID        int64
	Message   string
	CreatedAt pgtype.Timestamptz
}

func (q *Queries) CreateExample(ctx context.Context, arg CreateExampleParams) (Example, error) {
	row := q.db.QueryRow(ctx, createExample, arg.ID, arg.Message, arg.CreatedAt)
	var i Example
	err := row.Scan(&i.ID, &i.Message, &i.CreatedAt)
	return i, err
}

const deleteExample = `-- name: DeleteExample :exec
DELETE FROM examples
WHERE id = $1
`

func (q *Queries) DeleteExample(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteExample, id)
	return err
}

const getExample = `-- name: GetExample :one
SELECT id, message, created_at FROM examples
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetExample(ctx context.Context, id int64) (Example, error) {
	row := q.db.QueryRow(ctx, getExample, id)
	var i Example
	err := row.Scan(&i.ID, &i.Message, &i.CreatedAt)
	return i, err
}

const listExamples = `-- name: ListExamples :many
SELECT id, message, created_at FROM examples
`

func (q *Queries) ListExamples(ctx context.Context) ([]Example, error) {
	rows, err := q.db.Query(ctx, listExamples)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Example
	for rows.Next() {
		var i Example
		if err := rows.Scan(&i.ID, &i.Message, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateExample = `-- name: UpdateExample :exec
UPDATE examples
  set message = $2,
  created_at = $3
WHERE id = $1
`

type UpdateExampleParams struct {
	ID        int64
	Message   string
	CreatedAt pgtype.Timestamptz
}

func (q *Queries) UpdateExample(ctx context.Context, arg UpdateExampleParams) error {
	_, err := q.db.Exec(ctx, updateExample, arg.ID, arg.Message, arg.CreatedAt)
	return err
}
