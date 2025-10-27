package handler

import (
	model "booking/internal/app/repository"
	"booking/internal/app/services"
	"booking/pkgs/parser"
	"booking/pkgs/response"

	"net/http"
)

type ReqUserBooking struct {
	UserID int32 `json:"user_id"`
}

func ListBookingSeat(w http.ResponseWriter, r *http.Request) {
	var user_booking ReqUserBooking
	body, err := parser.Body(r.Body, user_booking)
	if err != nil {
		response.Error(w, err)
		return
	}
	data, err := service.ListBookingSeat(body.UserID)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func CreateBookingSeat(w http.ResponseWriter, r *http.Request) {
	var booking_seat model.CreateBookingSeatParams
	body, err := parser.Body(r.Body, booking_seat)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.CreateBookingSeat(body)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func DeleteBookingSeat(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/booking/seat/delete/")
	if err != nil {
		response.Error(w, err)
		return
	}
	if err := service.DeleteBookingSeat(id); err != nil {
		response.Error(w, err)
		return
	}
	response.Success(w, struct {
		Message string `json:"message"`
	}{Message: "booking deleted"})
}
