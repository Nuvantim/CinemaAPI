-- name: ListScreenType :many
SELECT * FROM screen_type;

-- name: CreateScreenType :one
INSERT INTO screen_type (name)
VALUES ($1)
RETURNING id;

-- name: GetScreenType :one
SELECT * FROM screen_type WHERE id = $1;

-- name: UpdateScreenType :one
UPDATE screen_type SET name = $2 WHERE id = $1 RETURNING id;

-- name: DeleteScreenType :exec
DELETE FROM screen_type WHERE id = $1;


-- name: ListScreen :many
SELECT
    screen.id,
    screen.name,
    cinema.name AS cinema_name,
    screen_type.name AS screen_type
FROM screen
INNER JOIN cinema ON screen.cinema_id = cinema.id
INNER JOIN screen_type ON screen.screen_type_id = screen_type.id;


-- name: CreateScreen :one
INSERT INTO screen (cinema_id, name, screen_type_id)
SELECT $1, $2, $3
WHERE EXISTS (SELECT 1 FROM cinema WHERE id = $1)
  AND EXISTS (SELECT 1 FROM screen_type WHERE id = $3)
RETURNING id;

-- name: GetScreen :one
SELECT
    screen.id,
    screen.name,
    cinema.name AS cinema_name,
    screen_type.name AS screen_type
FROM screen
INNER JOIN cinema ON screen.cinema_id = cinema.id
INNER JOIN screen_type ON screen.screen_type_id = screen_type.id
WHERE screen.id = $1;

-- name: UpdateScreen :one
UPDATE screen
SET cinema_id = $2,
    name = $3,
    screen_type_id = $4
WHERE screen.id = $1
  AND EXISTS (SELECT 1 FROM cinema WHERE id = $2)
  AND EXISTS (SELECT 1 FROM screen_type WHERE id = $4)
RETURNING id;

-- name: DeleteScreen :exec
DELETE FROM screen WHERE id = $1;

