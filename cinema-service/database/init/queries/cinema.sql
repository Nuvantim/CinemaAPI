-- name: ListCinema :many
SELECT * FROM cinema;

-- name: ListCinemaSchedule :many
SELECT
    c.name AS cinema_name,
    f.title AS film_title,
    st.name AS screen_type,
    sc.name AS screen_name,
    sh.start_time,
    sh.base_price
FROM showtime sh
INNER JOIN film f ON sh.film_id = f.id
INNER JOIN screen sc ON sh.screen_id = sc.id
INNER JOIN screen_type st ON sc.screen_type_id = st.id
INNER JOIN cinema c ON sc.cinema_id = c.id
WHERE c.id = $1
ORDER BY sh.start_time ASC;

-- name: CreateCinema :one
INSERT INTO cinema (name,address,city) VALUES ($1,$2,$3) RETURNING *;

-- name: GetCinema :one
SELECT * FROM cinema WHERE id = $1;

-- name: UpdateCinema :one
UPDATE cinema SET name=$2, address=$3, city=$4 WHERE id =$1 RETURNING id;

-- name: DeleteCinema :exec
DELETE FROM cinema WHERE id = $1;

