-- name: ListGenre :many
SELECT * FROM genre;

-- name: CreateGenre :one
INSERT INTO genre (name) VALUES ($1) RETURNING *;

-- name: GetGenre :one
SELECT * FROM genre WHERE id = $1;

-- name: UpdateGenre :one
UPDATE genre SET name =$2 WHERE id=$1 RETURNING id; 

-- name: DeleteGenre :exec
DELETE FROM genre WHERE id = $1;

