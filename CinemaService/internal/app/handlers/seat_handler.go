package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/services"
	"cinema/pkgs/helper"
	"net/http"
)

func ListSeat(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListSeat()
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func GetSeat(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/seat/")
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.GetSeat(id)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func CreateSeat(w http.ResponseWriter, r *http.Request) {
	var seat model.CreateSeatParams
	body, err := helper.ParserBody(r.Body, seat)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.CreateSeat(body)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func UpdateSeat(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/seat/update/")
	if err != nil {
		helper.Error(w, err)
	}

	var seat model.UpdateSeatParams
	body, err := helper.ParserBody(r.Body, seat)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.UpdateSeat(id, body)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)

}

func DeleteSeat(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/seat/delete/")
	if err != nil {
		helper.Error(w, err)
	}

	if err := service.DeleteSeat(id); err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, struct {
		Message string `json:"message"`
	}{Message: "seat deleted"})
}
