-- name: CreatePayment :one
INSERT INTO payment (booking_id, user_id,payment_method, payment_status, transaction_amount, payment_time)
VALUES(sqlc.arg(booking_id),sqlc.arg(user_id),sqlc.arg(payment_method),'Success' ,sqlc.arg(total_amount),NOW())
RETURNING *;

-- name: ListPayment :many
SELECT * FROM payment WHERE user_id = $1;
