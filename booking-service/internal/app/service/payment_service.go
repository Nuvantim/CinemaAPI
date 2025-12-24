package service

import (
	db "booking/database"
	model "booking/internal/app/repository"
	ctx "context"
	"fmt"
)

func ListPayment(user_id int64) ([]model.Payment, error) {
	data, err := db.Queries.ListPayment(ctx.Background(), user_id)
	if err != nil {
		return []model.Payment{}, db.Fatal(err)
	}
	return data, nil
}

func CreatePayment(body model.CreatePaymentParams) (model.Payment, error) {
	// get total ammount booking
	booking_paid, err := db.Queries.GetTotalAmmountBooking(ctx.Background(), body.BookingID)
	if err != nil {
		return model.Payment{}, db.Fatal(err)
	}
	if booking_paid == 0 {
		return model.Payment{}, fmt.Errorf("you haven't booked any seats yet")
	}

	// Create Payment
	data, err := db.Queries.CreatePayment(ctx.Background(), body)
	if err != nil {
		return model.Payment{}, db.Fatal(err)
	}
	// delete booking
	if err := db.Queries.DeleteBooking(ctx.Background(), body.BookingID); err != nil {
		return model.Payment{}, db.Fatal(err)
	}

	return data, nil
}
