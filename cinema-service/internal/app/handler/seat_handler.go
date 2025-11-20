package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/service"
	"cinema/pkg/parser"
	"cinema/pkg/response"
	"net/http"
)

func ListSeat(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListSeat()
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func GetSeat(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/seat/")
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.GetSeat(id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func CreateSeat(w http.ResponseWriter, r *http.Request) {
	var seat model.CreateSeatParams
	body, err := parser.Body(r.Body, seat)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.CreateSeat(body)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func SeatPrice(w http.ResponseWriter, r *http.Request) {
	var seat model.SeatPriceParams

	body, err := parser.Body(r.Body, seat)
	if err != nil {
		response.Error(w, err)
		return
	}
	data, err := service.SeatPrice(body)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Success(w, data)
}

func UpdateSeat(w http.ResponseWriter, r *http.Request) {
	var seat model.UpdateSeatParams
	body, err := parser.Body(r.Body, seat)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.UpdateSeat(body)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)

}

func DeleteSeat(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/seat/delete/")
	if err != nil {
		response.Error(w, err)
		return
	}

	if err := service.DeleteSeat(id); err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, struct {
		Message string `json:"message"`
	}{Message: "seat deleted"})
}
