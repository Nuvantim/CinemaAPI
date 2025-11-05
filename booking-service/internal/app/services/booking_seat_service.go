package service

import (
	db "booking/database"
	model "booking/internal/app/repository"
	ctx "context"

	"github.com/google/uuid"
)

func ListBookingSeat(booking_id int64) ([]model.BookingSeat, error) {

	data, err := db.Queries.ListBookingSeat(ctx.Background(), booking_id)
	if err != nil {
		return []model.BookingSeat{}, db.Fatal(err)
	}
	return data, nil
}

func CreateBookingSeat(body model.CreateBookingSeatParams) ([]model.BookingSeat, error) {
	booking_id, err := db.Queries.CreateBookingSeat(ctx.Background(), body)
	if err != nil {
		return []model.BookingSeat{}, db.Fatal(err)
	}

	// update booking amount
	if err := db.Queries.UpdateBookingAmount(ctx.Background(), body.BookingID); err != nil {
		return []model.BookingSeat{}, db.Fatal(err)
	}
	data, err := db.Queries.ListBookingSeat(ctx.Background(), booking_id)
	if err != nil {
		return []model.BookingSeat{}, db.Fatal(err)
	}
	return data, nil
}

func DeleteBookingSeat(id uuid.UUID) error {
	booking_id, err := db.Queries.DeleteBookingSeat(ctx.Background(), id)
	if err != nil {
		return db.Fatal(err)
	}
	// update booking amount
	if err := db.Queries.UpdateBookingAmount(ctx.Background(), booking_id); err != nil {
		return db.Fatal(err)
	}
	return nil
}
