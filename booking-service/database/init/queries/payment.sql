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
