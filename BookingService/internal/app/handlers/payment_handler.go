package handler

import(
	"booking/internal/app/services"
	"booking/pkgs/helper"
	model "booking/internal/app/repository"

	"net/http"
)

func ListPayment(w http.ResponseWriter, r *http.Request){
	var user_booking ReqUserBooking
	body,err := helper.ParserBody(r.Body, user_booking)
	if err != nil{
		helper.Error(w, err)
	}

	data,err := service.ListPayment(body.UserID)
	if err != nil{
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func CreatePayment(w http.ResponseWriter, r *http.Request){
	var payment model.CreatePaymentParams
	body,err := helper.ParserBody(r.Body, payment)
	if err != nil{
		helper.Error(w, err)
	}

	data,err := service.CreatePayment(body)
	if err != nil{
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func ReportProfit(w http.ResponseWriter, r *http.Request){
	data,err := service.ReportProfit()
	if err != nil{
		helper.Error(w,data)
	}
	helper.Success(w,data)
}