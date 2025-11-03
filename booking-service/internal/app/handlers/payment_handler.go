package handler

import (
	model "booking/internal/app/repository"
	"booking/internal/app/services"
	"booking/pkgs/parser"
	"booking/pkgs/response"

	"net/http"
)

func ListPayment(w http.ResponseWriter, r *http.Request) {
	var user_booking = struct {
		UserID int64 `json:"user_id"`
	}{}

	body, err := parser.Body(r.Body, user_booking)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.ListPayment(body.UserID)
	if err != nil {
		response.Error(w, err)
		return
	}

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

	response.Success(w, data)
}

// func ReportProfit(w http.ResponseWriter, r *http.Request) {
// 	data, err := service.ReportProfit()
// 	if err != nil {
// 		response.Error(w, err)
// 		return
// 	}
// 	response.Success(w, data)
// }
