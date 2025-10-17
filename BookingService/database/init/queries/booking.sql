-- name: CreateBooking :one
INSERT INTO booking (user_id, showtime_id, total_amount)
SELECT $1, $2, 0
WHERE EXISTS (SELECT id FROM user_account WHERE id = $1)
  AND EXISTS (SELECT id FROM showtime WHERE id = $2)
RETURNING id;

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
WHERE booking.id = $1;

-- name: UpdateBooking :one
UPDATE booking SET total_amount = (SELECT SUM(price_paid) FROM booking_seat WHERE booking_id = $1) 
WHERE id = $1 RETURNING total_amount;
