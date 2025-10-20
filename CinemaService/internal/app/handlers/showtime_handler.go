package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/services"
	"cinema/pkgs/helper"
	"net/http"
)

func ListShowTime(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListShowTime()
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func GetShowTime(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/showtime/")
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.GetShowTime(id)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func CreateShowTime(w http.ResponseWriter, r *http.Request) {
	var showtime model.CreateShowTimeParams
	body, err := helper.ParserBody(r.Body, showtime)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.CreateShowTime(body)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}

func UpdateShowTime(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/showtime/update/")
	if err != nil {
		helper.Error(w, err)
	}

	var showtime model.UpdateShowTimeParams
	body, err := helper.ParserBody(r.Body, showtime)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.UpdateShowTime(id, body)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)

}

func DeleteShowTime(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/showtime/delete/")
	if err != nil {
		helper.Error(w, err)
	}

	if err := service.DeleteShowTime(id); err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, struct {
		Message string `json:"message"`
	}{Message: "showtime deleted"})
}
