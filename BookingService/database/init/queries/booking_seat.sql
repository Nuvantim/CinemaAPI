-- name: CreateBookingSeat :one
INSERT INTO booking_seat (booking_id, seat_id, price_paid) SELECT $1,$2,$3
WHERE EXISTS (SELECT id FROM booking WHERE id = $1 AND user_id = $4) 
RETURNING booking_id;

-- name: ListBookingSeat :many
SELECT booking_seat.id, booking_seat.booking_id, seat.seat_row AS seat_row, seat.seat_number AS seat_number 
FROM booking_seat 
INNER JOIN seat ON booking_seat.seat_id = seat.id
WHERE booking_seat.booking_id = (SELECT id FROM booking WHERE user_id = $1);

