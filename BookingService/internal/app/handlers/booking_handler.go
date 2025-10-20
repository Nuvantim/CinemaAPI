package handler

import (
	"booking/internal/app/services"
	"booking/pkgs/helper"
	model "booking/internal/app/repository"

	"net/http"
)

func ListBooking(w http.ResponseWriter, r *http.Request){
	var booking model.Booking 

	body,err := helper.ParserBody(r.Body, booking)
	if err != nil{
		helper.Error(w, err)
	}
	data,err := service.ListBooking(body.UserID)
	if err != nil{
		helper.Error(w, err)
	}
	helper.Success(w, data)
}

func CreateBooking(w http.ResponseWriter, r *http.Request){
	var booking model.CreateBookingParams

	body,err := helper.ParserBody(r.Body, booking)
	if err != nil{
		helper.Error(w, err)
	}

	data,err := service.CreateBooking(body)
	if err != nil{
		helper.Error(w, err)
	}
	helper.Success(w, data)

}

func DeleteBooking(w http.ResponseWriter, r *http.Request){
	id,err := helper.ParserInt(r,"/booking/delete/")
	if err != nil{
		helper.Error(w,err)
	}
	if err := service.DeleteBooking(id);err != nil{
		helper.Error(w,err)
	}
	helper.Success(w, struct{Message string `json:"message"`}{Message : "booking deleted"})
}
