package service

import (
	db "booking/database"
	model "booking/internal/app/repository"
	ctx "context"
)

func ListBooking(id int32) ([]model.ListBookingRow, error) {
	data, err := db.Queries.ListBooking(ctx.Background(), id)
	if err != nil {
		return []model.ListBookingRow{}, db.Fatal(err)
	}

	return data, nil
}

func CreateBooking(body model.CreateBookingParams) ([]model.ListBookingRow, error) {
	booking_id, err := db.Queries.CreateBooking(ctx.Background(), body)
	if err != nil {
		return []model.ListBookingRow{}, db.Fatal(err)
	}

	data, err := db.Queries.ListBooking(ctx.Background(), booking_id)
	if err != nil {
		return []model.ListBookingRow{}, db.Fatal(err)
	}

	return data, nil
}

func DeleteBooking(id int32) error {
	if err := db.Queries.DeleteBooking(ctx.Background(), id); err != nil {
		return db.Fatal(err)
	}
	return nil
}
