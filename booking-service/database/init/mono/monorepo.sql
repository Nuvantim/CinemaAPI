-- name: CreateBooking :one
INSERT INTO booking (user_id, showtime_id)
VALUES($1, $2) RETURNING user_id;

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

-- name: CreateBookingSeat :one
INSERT INTO booking_seat (booking_id, seat_id, price_paid) VALUES($1,$2,$3) RETURNING *;

-- name: ListBookingSeat :many
SELECT * FROM booking_seat WHERE booking_id = $1;

-- name: DeleteBookingSeat :one
DELETE FROM booking_seat WHERE id = $1 RETURNING booking_id;

-- name: CreatePayment :one
INSERT INTO payment (booking_id, user_id, payment_method, payment_status, transaction_amount, payment_time)
SELECT 
    b.id,
    b.user_id,
    sqlc.arg(payment_method),
    sqlc.arg(payment_status),
    b.total_amount,
    NOW()
FROM booking b
WHERE b.id = sqlc.arg(booking_id)
RETURNING *;

-- name: ListPayment :many
SELECT * FROM payment WHERE user_id = $1 ORDER BY payment_time ASC;

