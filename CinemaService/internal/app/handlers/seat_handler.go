package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/services"
	"cinema/pkgs/parser"
	"cinema/pkgs/response"
	"net/http"
)

func ListSeat(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListSeat()
	if err != nil {
		response.Error(w, err)
	}

	response.Success(w, data)
}

func GetSeat(w http.ResponseWriter, r *http.Request) {
	id, err := parser.Params(r, "/seat/")
	if err != nil {
		response.Error(w, err)
	}

	data, err := service.GetSeat(id)
	if err != nil {
		response.Error(w, err)
	}

	response.Success(w, data)
}

func CreateSeat(w http.ResponseWriter, r *http.Request) {
	var seat model.CreateSeatParams
	body, err := parser.Body(r.Body, seat)
	if err != nil {
		response.Error(w, err)
	}

	data, err := service.CreateSeat(body)
	if err != nil {
		response.Error(w, err)
	}

	response.Success(w, data)
}

func UpdateSeat(w http.ResponseWriter, r *http.Request) {
	id, err := parser.Params(r, "/seat/update/")
	if err != nil {
		response.Error(w, err)
	}

	var seat model.UpdateSeatParams
	body, err := parser.Body(r.Body, seat)
	if err != nil {
		response.Error(w, err)
	}

	data, err := service.UpdateSeat(id, body)
	if err != nil {
		response.Error(w, err)
	}

	response.Success(w, data)

}

func DeleteSeat(w http.ResponseWriter, r *http.Request) {
	id, err := parser.Params(r, "/seat/delete/")
	if err != nil {
		response.Error(w, err)
	}

	if err := service.DeleteSeat(id); err != nil {
		response.Error(w, err)
	}

	response.Success(w, struct {
		Message string `json:"message"`
	}{Message: "seat deleted"})
}
