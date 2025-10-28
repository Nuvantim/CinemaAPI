-- name: ListCinema :many
SELECT * FROM cinema;

-- name: CreateCinema :one
INSERT INTO cinema (name,address,city) VALUES ($1,$2,$3) RETURNING *;

-- name: GetCinema :one
SELECT * FROM cinema WHERE id = $1;

-- name: UpdateCinema :one
UPDATE cinema SET name=$2, address=$3, city=$4 WHERE id =$1 RETURNING id;

-- name: DeleteCinema :exec
DELETE FROM cinema WHERE id = $1;
