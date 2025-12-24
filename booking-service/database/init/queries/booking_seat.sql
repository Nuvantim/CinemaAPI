-- name: CreateBookingSeat :one
INSERT INTO booking_seat (booking_id, seat_id, price_paid) VALUES($1,$2,$3) RETURNING *;

-- name: ListBookingSeat :many
SELECT * FROM booking_seat WHERE booking_id = $1;

-- name: DeleteBookingSeat :one
DELETE FROM booking_seat WHERE id = $1 RETURNING booking_id;

