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

-- name: ListFilm :many
SELECT film.id, film.title, film.director, genre.name AS genre, film.duration FROM film 
INNER JOIN genre ON film.genre_id = genre.id ORDER BY film.id ASC;

-- name: CreateFilm :one
INSERT INTO film (title,director,genre_id,duration) SELECT $1,$2,$3,$4 
WHERE EXISTS(SELECT id FROM genre WHERE id = $3) RETURNING id;

-- name: GetFilm :one
SELECT film.id, film.title, film.director, genre.name AS genre, film.duration FROM film  
INNER JOIN genre  ON film.genre_id = genre.id WHERE film.id=$1;

-- name: SearchFilm :many
SELECT film.id, film.title, film.director, genre.name AS genre, film.duration FROM film
INNER JOIN genre  ON film.genre_id = genre.id
WHERE film.title LIKE $1;

-- name: SearchGenreFilm :many
SELECT film.id, film.title, film.director, genre.name AS genre, film.duration FROM film  
INNER JOIN genre  ON film.genre_id = genre.id WHERE film.genre_id=(SELECT id FROM genre WHERE genre.id = $1);

-- name: UpdateFilm :one
UPDATE film SET title=$2, director=$3, genre_id = $4, duration=$5 WHERE film.id = $1 
AND EXISTS(SELECT id FROM genre WHERE genre.id = $4) RETURNING film.id;

-- name: DeleteFilm :exec
DELETE FROM film WHERE id = $1;



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

-- name: ListScreenType :many
SELECT * FROM screen_type;

-- name: CreateScreenType :one
INSERT INTO screen_type (name)
VALUES ($1)
RETURNING id;

-- name: GetScreenType :one
SELECT * FROM screen_type WHERE id = $1;

-- name: UpdateScreenType :one
UPDATE screen_type
SET name = $2
WHERE id = $1
RETURNING id;

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

-- name: ListSeat :many
SELECT
    seat.id,
    screen.name AS screen_name,
    seat.seat_row,
    seat.seat_number,
    seat.seat_price_modifier
FROM seat
INNER JOIN screen ON seat.screen_id = screen.id;

-- name: GetSeat :one
SELECT
    seat.id,
    screen.name AS screen_name,
    seat.seat_row,
    seat.seat_number,
    seat.seat_price_modifier
FROM seat
INNER JOIN screen ON seat.screen_id = screen.id
WHERE seat.id = $1;

-- name: CreateSeat :one
INSERT INTO seat (screen_id, seat_row, seat_number, seat_price_modifier)
SELECT $1, $2, $3, $4
WHERE EXISTS (SELECT 1 FROM screen WHERE id = $1)
RETURNING id;

-- name: UpdateSeat :one
UPDATE seat
SET screen_id = $2,
    seat_row = $3,
    seat_number = $4,
    seat_price_modifier = $5
WHERE seat.id = $1
  AND EXISTS (SELECT 1 FROM screen WHERE id = $2)
RETURNING id;

-- name: SeatPrice :one
SELECT (showtime.base_price * seat.seat_price_modifier) AS seat_price 
FROM showtime JOIN seat ON showtime.screen_id = seat.screen_id 
WHERE showtime.id = sqlc.arg(showtime_id) AND seat.id = sqlc.arg(seat_id);

-- name: DeleteSeat :exec
DELETE FROM seat WHERE id = $1;

