package handler

import (
	model "booking/internal/app/repository"
	"booking/internal/app/service"
	"booking/pkgs/parser"
	"booking/pkgs/response"

	"net/http"
)

func ListBookingSeat(w http.ResponseWriter, r *http.Request) {
	var user_booking = struct {
		BookingID int64 `json:"booking_id"`
	}{}
	body, err := parser.Body(r.Body, user_booking)
	if err != nil {
		response.Error(w, err)
		return
	}
	data, err := service.ListBookingSeat(body.BookingID)
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
	id, err := parser.ParamsUUID(r, "/booking/seat/delete/")
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
