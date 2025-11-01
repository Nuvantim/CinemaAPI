-- name: CreatePayment :one
INSERT INTO payment (booking_id, payment_method, payment_status, transaction_amount, payment_time)
SELECT 
    b.id,
    $2 AS payment_method,
    "Success" AS payment_status,
    b.total_amount AS transaction_amount,
    NOW() AS payment_time
FROM booking
WHERE booking.id = $1
RETURNING *;

-- name: ListPayment :many
SELECT * FROM payment WHERE booking_id = (SELECT id FROM booking WHERE user_id = $1);