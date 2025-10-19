package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/services"
	"cinema/pkgs/helper"
	"net/http"
)

func ListCinema(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListCinema()
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func GetCinema(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/film/")
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.GetCinema(id)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func CreateCinema(w http.ResponseWriter, r *http.Request) {
	var cinema model.CreateCinemaParams
	body, err := helper.ParserBody(r.Body, cinema)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.CreateCinema(body)
	if err != nil {
		helper.Error(w, err)
	}
	helper.Success(w, data)
}

func UpdateCinema(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/cinema/update/")
	if err != nil {
		helper.Error(w, err)
	}

	var cinema model.UpdateCinemaParams
	body, err := helper.ParserBody(r.Body, cinema)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.UpdateCinema(id, body)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func DeleteCinema(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/film/delete")
	if err != nil {
		helper.Error(w, err)
	}

	if err := service.DeleteFilm(id); err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, struct {
		Message string `json:"message"`
	}{Message: "cinema deleted"})
}
