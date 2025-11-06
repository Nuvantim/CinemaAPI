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
	booking_paid,err := db.Queries.GetTotalAmmountBooking(ctx.Background(), body.BookingID)
	if err != nil{
		return model.Payment{},db.Fatal(err)
	}
	if booking_paid.Float64 == 0 {
		return model.Payment{},fmt.Errorf("you haven't booked any seats yet")
	}
	
	// Create Payment
	data, err := db.Queries.CreatePayment(ctx.Background(), body)
	if err != nil {
		return model.Payment{}, db.Fatal(err)
	}
	// delete booking
	if err := db.Queries.DeleteBooking(ctx.Background(), body.BookingID);err != nil{
		return model.Payment{}, db.Fatal(err)
	}

	return data, nil
}

// func ReportProfit() ([]model.ReportProfitRow, error) {
// 	data, err := db.Queries.ReportProfit(ctx.Background())
// 	if err != nil {
// 		return []model.ReportProfitRow{}, db.Fatal(err)
// 	}
// 	return data, nil
// }

// result, err := db.Queries.UpdateBooking(ctx.Background(), body.BookingID)
// 	if err != nil {
// 		return model.Payment{}, db.Fatal(err)
// 	}

// 	if !result.Valid {
// 		return model.Payment{}, fmt.Errorf("invalid amount value")
// 	}

// 	// create payment
// 	body.TransactionAmount = result.Float64
