package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/service"
	"cinema/pkg/parser"
	"cinema/pkg/response"
	"net/http"
)

func ListScreenType(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListScreenType()
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Success(w, data)
}

func GetScreenType(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/screen/type/")
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.GetScreenType(id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func CreateScreenType(w http.ResponseWriter, r *http.Request) {
	var screen_type model.ScreenType
	body, err := parser.Body(r.Body, screen_type)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.CreateScreenType(body.Name)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func UpdateScreenType(w http.ResponseWriter, r *http.Request) {
	var screen_type model.UpdateScreenTypeParams
	body, err := parser.Body(r.Body, screen_type)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.UpdateScreenType(body)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func DeleteScreenType(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/screen/type/delete/")
	if err != nil {
		response.Error(w, err)
		return
	}

	if err := service.DeleteScreenType(id); err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, struct {
		Message string `json:"message"`
	}{Message: "screen type deleted"})
}
