package service

import (
	db "booking/database"
	model "booking/internal/app/repository"
	ctx "context"
	"fmt"
)

func ListPayment(user_id int32) ([]model.Payment, error) {
	data, err := db.Queries.ListPayment(ctx.Background(), user_id)
	if err != nil {
		return []model.Payment{}, db.Fatal(err)
	}
	return data, nil
}

func CreatePayment(body model.CreatePaymentParams) (model.Payment, error) {
	// get amount booking
	result, err := db.Queries.UpdateBooking(ctx.Background(), body.BookingID)
	if err != nil {
		return model.Payment{}, db.Fatal(err)
	}

	if !result.Valid {
		return model.Payment{}, fmt.Errorf("invalid amount value")
	}

	// create payment
	body.TransactionAmount = result.Float64
	data, err := db.Queries.CreatePayment(ctx.Background(), body)
	if err != nil {
		return model.Payment{}, db.Fatal(err)
	}

	return data, nil
}

func ReportProfit() ([]model.ReportProfitRow, error) {
	data, err := db.Queries.ReportProfit(ctx.Background())
	if err != nil {
		return []model.ReportProfitRow{}, db.Fatal(err)
	}
	return data, nil
}
