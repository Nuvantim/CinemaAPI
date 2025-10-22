package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/services"
	"cinema/pkgs/parser"
	"cinema/pkgs/response"
	"net/http"
)

func ListScreen(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListScreen()
	if err != nil {
		response.Error(w, err)
	}
	response.Success(w, data)
}

func GetScreen(w http.ResponseWriter, r *http.Request) {
	id, err := parser.Params(r, "/screen/")
	if err != nil {
		response.Error(w, err)
	}

	data, err := service.GetScreen(id)
	if err != nil {
		response.Error(w, err)
	}

	response.Success(w, data)
}

func CreateScreen(w http.ResponseWriter, r *http.Request) {
	var screen model.CreateScreenParams
	body, err := parser.Body(r.Body, screen)
	if err != nil {
		response.Error(w, err)
	}

	data, err := service.CreateScreen(body)
	if err != nil {
		response.Error(w, err)
	}

	response.Success(w, data)
}

func UpdateScreen(w http.ResponseWriter, r *http.Request) {
	id, err := parser.Params(r, "/screen/update/")
	if err != nil {
		response.Error(w, err)
	}

	var screen model.UpdateScreenParams
	body, err := parser.Body(r.Body, screen)
	if err != nil {
		response.Error(w, err)
	}

	data, err := service.UpdateScreen(id, body)
	if err != nil {
		response.Error(w, err)
	}

	response.Success(w, data)

}

func DeleteScreen(w http.ResponseWriter, r *http.Request) {
	id, err := parser.Params(r, "/screen/delete/")
	if err != nil {
		response.Error(w, err)
	}

	if err := service.DeleteScreen(id); err != nil {
		response.Error(w, err)
	}

	response.Success(w, struct {
		Message string `json:"message"`
	}{Message: "screen deleted"})
}
