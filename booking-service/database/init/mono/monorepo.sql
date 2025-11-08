-- name: CreateBooking :one
INSERT INTO booking (user_id, showtime_id)
VALUES($1,$2) RETURNING *;

-- name: ListBooking :many
SELECT * FROM booking WHERE user_id = $1;

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
INSERT INTO payment (user_id,booking_id, payment_method, payment_status, transaction_amount, payment_time)
SELECT 
    sqlc.arg(user_id) AS user_id,
    b.id,
    sqlc.arg(payment_method) AS payment_method,
    'Success' AS payment_status,
    b.total_amount AS transaction_amount,
    NOW() AS payment_time
FROM booking b
WHERE b.id = sqlc.arg(booking_id)
RETURNING *;

-- name: ListPayment :many
SELECT * FROM payment WHERE user_id = $1;
