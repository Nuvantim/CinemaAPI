package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/services"
	"cinema/pkgs/helper"
	"net/http"
)

func ListScreen(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListScreen()
	if err != nil {
		helper.Error(w, err)
	}
	helper.Success(w, data)
}

func GetScreen(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/screen/")
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.GetScreen(id)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func CreateScreen(w http.ResponseWriter, r *http.Request) {
	var screen model.CreateScreenParams
	body, err := helper.ParserBody(r.Body, screen)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.CreateScreen(body)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func UpdateScreen(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/screen/update/")
	if err != nil {
		helper.Error(w, err)
	}

	var screen model.UpdateScreenParams
	body, err := helper.ParserBody(r.Body, screen)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.UpdateScreen(id, body)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)

}

func DeleteScreen(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/screen/delete/")
	if err != nil {
		helper.Error(w, err)
	}

	if err := service.DeleteScreen(id); err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, struct {
		Message string `json:"message"`
	}{Message: "screen deleted"})
}
