package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/services"
	"cinema/pkgs/parser"
	"cinema/pkgs/response"
	"net/http"
)

func ListCinema(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListCinema()
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func GetCinema(w http.ResponseWriter, r *http.Request) {
	id, err := parser.Params(r, "/cinema/")
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.GetCinema(id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func CreateCinema(w http.ResponseWriter, r *http.Request) {
	var cinema model.CreateCinemaParams
	body, err := parser.Body(r.Body, cinema)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.CreateCinema(body)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Success(w, data)
}

func UpdateCinema(w http.ResponseWriter, r *http.Request) {
	id, err := parser.Params(r, "/cinema/update/")
	if err != nil {
		response.Error(w, err)
		return
	}

	var cinema model.UpdateCinemaParams
	body, err := parser.Body(r.Body, cinema)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.UpdateCinema(id, body)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func DeleteCinema(w http.ResponseWriter, r *http.Request) {
	id, err := parser.Params(r, "/cinema/delete/")
	if err != nil {
		response.Error(w, err)
		return
	}

	if err := service.DeleteCinema(id); err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, struct {
		Message string `json:"message"`
	}{Message: "cinema deleted"})
}
