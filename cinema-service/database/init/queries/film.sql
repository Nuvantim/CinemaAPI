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



