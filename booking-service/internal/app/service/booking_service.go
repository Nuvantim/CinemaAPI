package service

import (
	db "booking/database"
	model "booking/internal/app/repository"
	ctx "context"

	crand "crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func ListBooking(id int64) ([]model.Booking, error) {
	data, err := db.Queries.ListBooking(ctx.Background(), id)
	if err != nil {
		return []model.Booking{}, db.Fatal(err)
	}

	return data, nil
}

func CreateBooking(body model.CreateBookingParams) (model.Booking, error) {
	id, err := func() (int64, error) {
		now := time.Now()
		str := now.Format("20060102150405")

		var id int64
		if _, err := fmt.Sscan(str, &id); err != nil {
			return 0, fmt.Errorf("invalid id format: %w", err)
		}

		number, _ := crand.Int(crand.Reader, big.NewInt(90))
		random := 10 + number.Int64()
		result := id*100 + random

		return result, nil
	}()

	if err != nil {
		return model.Booking{}, err
	}

	body.ID = id
	data, err := db.Queries.CreateBooking(ctx.Background(), body)
	if err != nil {
		return model.Booking{}, db.Fatal(err)
	}

	return data, nil
}

func GetBooking(id int64) (model.Booking, error) {
	data, err := db.Queries.GetBooking(ctx.Background(), id)
	if err != nil {
		return model.Booking{}, db.Fatal(err)
	}

	return data, nil
}

func DeleteBooking(id int64) error {
	if err := db.Queries.DeleteBooking(ctx.Background(), id); err != nil {
		return db.Fatal(err)
	}
	return nil
}
