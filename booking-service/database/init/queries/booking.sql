-- name: CreateBooking :one
INSERT INTO booking (id,user_id, showtime_id, total_amount)
VALUES($1, $2, $3, 0) RETURNING user_id;

-- name: ListBooking :many
SELECT * FROM booking WHERE user_id = $1;

-- name: UpdateBookingAmount :exec
UPDATE booking SET total_amount = (SELECT SUM(price_paid) FROM booking_seat WHERE booking_id = $1) 
WHERE id = $1;

-- name: DeleteBooking :exec
DELETE FROM booking WHERE id = $1;
