-- name: CreatePayment :one
INSERT INTO payment (booking_id, user_id,payment_method, payment_status, transaction_amount, payment_time)
SELECT 
    b.id,
    sqlc.arg(user_id) AS user_id,
    sqlc.arg(payment_method) AS payment_method,
    'Success' AS payment_status,
    b.total_amount AS transaction_amount,
    NOW() AS payment_time
FROM booking b
WHERE b.id = sqlc.arg(booking_id)
RETURNING *;

-- name: ListPayment :many
SELECT * FROM payment WHERE user_id = $1;
