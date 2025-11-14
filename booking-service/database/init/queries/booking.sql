-- name: CreateBooking :one
INSERT INTO booking (id,user_id, showtime_id)
VALUES($1, $2, $3) RETURNING *;

-- name: ListBooking :many
SELECT * FROM booking WHERE user_id = $1 ORDER BY booking_time ASC;

-- name: GetBooking :one
SELECT * FROM booking WHERE id =$1;

-- name: GetTotalAmmountBooking :one
SELECT total_amount FROM booking WHERE id = $1;

-- name: UpdateBookingAmount :exec
UPDATE booking SET total_amount = (SELECT SUM(price_paid) FROM booking_seat WHERE booking_id = $1) 
WHERE id = $1;

-- name: DeleteBooking :exec
DELETE FROM booking WHERE id = $1;
