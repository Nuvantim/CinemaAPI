-- name: CreatePayment :one
INSERT INTO payment (booking_id, payment_method, payment_status, transaction_amount, payment_time)
SELECT 
    b.id,
    $2 AS payment_method,
    "Success" AS payment_status,
    b.total_amount AS transaction_amount,
    NOW() AS payment_time
FROM booking b
WHERE b.id = $1
RETURNING *;

-- name: ListPayment :many
SELECT * FROM payment WHERE booking_id = (SELECT id FROM booking WHERE user_id = $1);

-- name: ReportProfit :many
SELECT
    F.title AS Film,
    C.city AS Kota,
    COUNT(B.id) AS Total_Tiket_Terjual,
    SUM(P.transaction_amount) AS Total_Pendapatan
FROM
    booking B
JOIN
    payment P ON B.id = P.booking_id
JOIN
    showtime ST ON B.showtime_id = ST.id
JOIN
    film F ON ST.film_id = F.id
JOIN
    screen S ON ST.screen_id = S.id
JOIN
    cinema C ON S.cinema_id = C.id
WHERE
    P.payment_status = 'Success'
GROUP BY
    F.title, C.city
ORDER BY
    Total_Pendapatan DESC;
