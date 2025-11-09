-- name: ListSeat :many
SELECT
    seat.id,
    screen.name AS screen_name,
    seat.seat_row,
    seat.seat_number,
    seat.seat_price_modifier
FROM seat
INNER JOIN screen ON seat.screen_id = screen.id;

-- name: GetSeat :one
SELECT
    seat.id,
    screen.name AS screen_name,
    seat.seat_row,
    seat.seat_number,
    seat.seat_price_modifier
FROM seat
INNER JOIN screen ON seat.screen_id = screen.id
WHERE seat.id = $1;

-- name: CreateSeat :one
INSERT INTO seat (screen_id, seat_row, seat_number, seat_price_modifier)
SELECT $1, $2, $3, $4
WHERE EXISTS (SELECT 1 FROM screen WHERE id = $1)
RETURNING id;

-- name: UpdateSeat :one
UPDATE seat
SET screen_id = $2,
    seat_row = $3,
    seat_number = $4,
    seat_price_modifier = $5
WHERE seat.id = $1
  AND EXISTS (SELECT 1 FROM screen WHERE id = $2)
RETURNING id;

-- name: SeatPrice :one
SELECT (showtime.base_price * seat.seat_price_modifier) AS seat_price 
FROM showtime JOIN seat ON showtime.screen_id = seat.screen_id 
WHERE showtime.id = sqlc.arg(showtime_id) AND seat.id = sqlc.arg(seat_id);

-- name: DeleteSeat :exec
DELETE FROM seat WHERE id = $1;

