-- name: CreateBooking :one
INSERT INTO booking (user_id, showtime_id, total_amount)
SELECT $1, $2, 0
WHERE EXISTS (SELECT id FROM user_account WHERE id = $1)
  AND EXISTS (SELECT id FROM showtime WHERE id = $2)
RETURNING user_id;

-- name: ListBooking :many
SELECT
    booking.id,
    user_account.name AS user_name,
    film.title AS film_name,
    showtime.start_time,
    booking.booking_time,
    booking.total_amount
FROM booking
INNER JOIN user_account ON booking.user_id = user_account.id
INNER JOIN showtime ON booking.showtime_id = showtime.id
INNER JOIN film ON showtime.film_id = film.id
WHERE user_id = $1;

-- name: UpdateBookingAmount :exec
UPDATE booking SET total_amount = (SELECT SUM(price_paid) FROM booking_seat WHERE booking_id = $1) 
WHERE id = $1;

-- name: DeleteBooking :exec
DELETE FROM booking WHERE id = $1;
-- name: CreateBookingSeat :one
INSERT INTO booking_seat (booking_id, seat_id, price_paid) SELECT $1,$2,$3
WHERE EXISTS (SELECT id FROM booking WHERE id = $1 AND user_id = $4) 
RETURNING booking_id;

-- name: ListBookingSeat :many
SELECT booking_seat.id, booking_seat.booking_id, seat.seat_row AS seat_row, seat.seat_number AS seat_number 
FROM booking_seat 
INNER JOIN seat ON booking_seat.seat_id = seat.id
WHERE booking_seat.booking_id = (SELECT id FROM booking WHERE user_id = $1);

-- name: DeleteBookingSeat :exec
DELETE FROM booking_seat WHERE id = $1;

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
