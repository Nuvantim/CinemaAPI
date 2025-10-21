package handler

import(
	"booking/internal/app/services"
	"booking/pkgs/helper"
	model "booking/internal/app/repository"

	"net/http"
)

type ReqUserBooking struct{
	UserID int32 `json:"user_id"`
}

func ListBookingSeat(w http.ResponseWriter, r *http.Request){
	var user_booking ReqUserBooking
	body,err := helper.ParserBody(r.Body, user_booking)
	if err != nil{
		helper.Error(w, err)
	}
	data,err := service.ListBookingSeat(body.UserID)
	if err != nil{
		helper.Error(w, err)
	}

	helper.Success(w,data)
}

func CreateBookingSeat(w http.ResponseWriter, r *http.Request){
	var booking_seat model.CreateBookingSeatParams
	body,err := helper.ParserBody(r.Body, booking_seat)
	if err != nil{
		helper.Error(w, err)
	}

	data,err := service.CreateBookingSeat(body)
	if err != nil{
		helper.Error(w, err)
	}

	helper.Success(w, data)
}