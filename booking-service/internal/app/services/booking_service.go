package service

import (
	db "booking/database"
	model "booking/internal/app/repository"
	ctx "context"

	"fmt"
	"time"
	"log"
)

func ListBooking(id int32) ([]model.Booking, error) {
	data, err := db.Queries.ListBooking(ctx.Background(), id)
	if err != nil {
		return []model.Booking{}, db.Fatal(err)
	}

	return data, nil
}

func CreateBooking(body model.CreateBookingParams) ([]model.Booking, error) {
	booking_id := GenBookingID()
	log.Println(booking_id)
	body.ID = booking_id
	booking_id, err := db.Queries.CreateBooking(ctx.Background(), body)
	if err != nil {
		return []model.Booking{}, db.Fatal(err)
	}

	data, err := db.Queries.ListBooking(ctx.Background(), booking_id)
	if err != nil {
		return []model.Booking{}, db.Fatal(err)
	}

	return data, nil
}

func DeleteBooking(id int32) error {
	if err := db.Queries.DeleteBooking(ctx.Background(), id); err != nil {
		return db.Fatal(err)
	}
	return nil
}

func GenBookingID() int64 {
	now := time.Now()
	str := now.Format("20060102150405")

	var id int64
	fmt.Sscan(str, &id)

	return id
}
