-- name: GetExample :one
SELECT * FROM examples
WHERE id = $1 LIMIT 1;

-- name: ListExamples :many
SELECT * FROM examples;

-- name: CreateExample :one
INSERT INTO examples (
  id, message, created_at
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateExample :exec
UPDATE examples
  set message = $2,
  created_at = $3
WHERE id = $1;

-- name: DeleteExample :exec
DELETE FROM examples
WHERE id = $1;