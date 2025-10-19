package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/services"
	"cinema/pkgs/helper"
	"net/http"
)

func ListScreenType(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListScreenType()
	if err != nil {
		helper.Error(w, err)
	}
	helper.Success(w, data)
}

func GetScreenType(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/screen/type/")
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.GetScreenType(id)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func CreateScreenType(w http.ResponseWriter, r *http.Request) {
	var screen_type model.ScreenType
	body, err := helper.ParserBody(r.Body, screen_type)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.CreateScreenType(body.Name)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func UpdateScreenType(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/screen/type/update/")
	if err != nil {
		helper.Error(w, err)
	}

	var screen_type model.UpdateScreenTypeParams
	body, err := helper.ParserBody(r.Body, screen_type)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.UpdateScreenType(id, body)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func DeleteScreenType(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/scree/type/delete/")
	if err != nil {
		helper.Error(w, err)
	}

	if err := service.DeleteScreenType(id); err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, struct {
		Message string `json:"message"`
	}{Message: "screen type deleted"})
}
