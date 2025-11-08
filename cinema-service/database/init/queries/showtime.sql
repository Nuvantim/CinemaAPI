-- name: ListShowTime :many
SELECT
    showtime.id,
    film.title AS film,
    screen.name AS screen,
    showtime.start_time,
    showtime.base_price
FROM showtime
INNER JOIN film ON showtime.film_id = film.id
INNER JOIN screen ON showtime.screen_id = screen.id;

-- name: GetShowTime :one
SELECT
    showtime.id,
    film.title AS film,
    screen.name AS screen,
    showtime.start_time,
    showtime.base_price
FROM showtime
INNER JOIN film ON showtime.film_id = film.id
INNER JOIN screen ON showtime.screen_id = screen.id
WHERE showtime.id = $1;

-- name: CreateShowTime :one
INSERT INTO showtime (film_id, screen_id, start_time, base_price)
SELECT $1, $2, $3, $4
WHERE EXISTS (SELECT 1 FROM film WHERE id = $1)
  AND EXISTS (SELECT 1 FROM screen WHERE id = $2)
RETURNING id;

-- name: UpdateShowTime :one
UPDATE showtime
SET
    film_id = $2,
    screen_id = $3,
    start_time = $4,
    base_price = $5
WHERE showtime.id = $1
  AND EXISTS (SELECT 1 FROM film WHERE id = $2)
  AND EXISTS (SELECT 1 FROM screen WHERE id = $3)
RETURNING id;

-- name: DeleteShowTime :exec
DELETE FROM showtime WHERE id = $1;


