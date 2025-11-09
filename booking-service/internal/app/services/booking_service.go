package service

import (
	db "booking/database"
	model "booking/internal/app/repository"
	ctx "context"

	"fmt"
	"math/rand"
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
	id := GenBookingId()

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

func GenBookingId() int64 {
	now := time.Now()
	str := now.Format("20060102150405")

	var id int64
	fmt.Sscan(str, &id)

	random := int64(rand.Intn(90) + 10)
	result := id*100 + random

	return result
}
