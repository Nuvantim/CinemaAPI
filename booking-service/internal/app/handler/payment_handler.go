package handler

import (
	model "booking/internal/app/repository"
	"booking/internal/app/service"
	"booking/pkg/parser"
	"booking/pkg/response"
	rds "booking/redis"

	"fmt"
	"net/http"
)

func ListPayment(w http.ResponseWriter, r *http.Request) {
	var user = struct {
		UserID int64 `json:"user_id"`
	}{}

	body, err := parser.Body(r.Body, user)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.ListPayment(body.UserID)
	if err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		_ = rds.SetData(fmt.Sprintf("list:payment:%d", user.UserID), data)
	}()

	response.Success(w, data)
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment model.CreatePaymentParams
	body, err := parser.Body(r.Body, payment)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.CreatePayment(body)
	if err != nil {
		response.Error(w, err)
		return
	}
	var userId int64 = data.UserID

	// set data to redis
	go func(userId int64) {
		// set data payment
		data_payment, _ := service.ListPayment(userId)
		if data_payment != nil {
			_ = rds.SetData(fmt.Sprintf("list:payment:%d", userId), data_payment)
		}

		// set data booking
		data_booking, _ := service.ListBooking(userId)
		_ = rds.SetData(fmt.Sprintf("list:booking:%d", userId), data_booking)
	}(userId)

	response.Success(w, data)
}
