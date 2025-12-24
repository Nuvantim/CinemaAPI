package service

import (
	db "booking/database"
	model "booking/internal/app/repository"
	rds "booking/redis"
	ctx "context"

	"fmt"

	"github.com/google/uuid"
)

func ListBookingSeat(booking_id int64) ([]model.BookingSeat, error) {

	data, err := db.Queries.ListBookingSeat(ctx.Background(), booking_id)
	if err != nil {
		return []model.BookingSeat{}, db.Fatal(err)
	}
	return data, nil
}

func CreateBookingSeat(body model.CreateBookingSeatParams) (model.BookingSeat, error) {
	data, err := db.Queries.CreateBookingSeat(ctx.Background(), body)
	if err != nil {
		return model.BookingSeat{}, db.Fatal(err)
	}

	// update booking amount
	if err := db.Queries.UpdateBookingAmount(ctx.Background(), body.BookingID); err != nil {
		return model.BookingSeat{}, db.Fatal(err)
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
	// get data booking
	booking, _ := GetBooking(booking_id)
	var user_id int64 = booking.UserID

	// set data to redis
	go func(userId int64) {
		data_booking, _ := ListBooking(userId)
		_ = rds.SetData(fmt.Sprintf("list:booking:%d", userId), data_booking)
	}(user_id)

	return nil
}
