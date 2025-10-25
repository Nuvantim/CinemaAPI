package handler

import (
	model "booking/internal/app/repository"
	"booking/internal/app/services"
	"booking/pkgs/parser"
	"booking/pkgs/response"

	"net/http"
)

func ListBooking(w http.ResponseWriter, r *http.Request) {
	var booking model.Booking

	body, err := parser.Body(r.Body, booking)
	if err != nil {
		response.Error(w, err)
		return
	}
	data, err := service.ListBooking(body.UserID)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Success(w, data)
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking model.CreateBookingParams

	body, err := parser.Body(r.Body, booking)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.CreateBooking(body)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Success(w, data)

}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	id, err := parser.Params(r, "/booking/delete/")
	if err != nil {
		response.Error(w, err)
		return
	}
	if err := service.DeleteBooking(id); err != nil {
		response.Error(w, err)
		return
	}
	response.Success(w, struct {
		Message string `json:"message"`
	}{Message: "booking deleted"})
}
