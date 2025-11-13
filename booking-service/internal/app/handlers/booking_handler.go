package handler

import (
	model "booking/internal/app/repository"
	"booking/internal/app/services"
	"booking/pkgs/parser"
	"booking/pkgs/response"
	rds "booking/redis"

	"fmt"
	"net/http"
)

func ListBooking(w http.ResponseWriter, r *http.Request) {
	var booking = struct {
		UserID int64 `json:"user_id"`
	}{}

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

	// set data to redis
	defer func() {
		_ := rds.SetData(fmt.Sprintf("list:booking:%d", body.UserId), data)
	}()

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

	// set data to redis
	defer func() {
		data_booking, _ := service.ListBooking(data.UserID)
		if data != nil {
			_ := rds.SetData(fmt.Sprintf("list:booking:%d", data.UserId), data_booking)
		}
	}()

	response.Success(w, data)

}

func GetBooking(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/booking/")
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.GetBooking(id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/booking/delete/")
	if err != nil {
		response.Error(w, err)
		return
	}

	booking, _ := service.GetBooking(id)

	if err := service.DeleteBooking(id); err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	defer func() {
		if booking.UserID != 0 {
			data_booking, _ := service.ListBooking(booking.UserID)
			_ = rds.SetData(fmt.Sprintf("list:booking:%d", dataBooking.UserID), data_booking)
		}
	}()

	response.Success(w, struct {
		Message string `json:"message"`
	}{Message: "booking deleted"})
}

