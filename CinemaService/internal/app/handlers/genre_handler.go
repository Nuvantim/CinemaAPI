package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/services"
	"cinema/pkgs/helper"
	"net/http"
)

func ListGenre(w http.ResponseWriter, r *http.Request) {
	genre, err := service.ListGenre()
	if err != nil {
		helper.Error(w, err)
	}
	helper.Success(w, genre)
}

func GetGenre(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/genre/")
	if err != nil {
		helper.Error(w, err)
	}

	genre, err := service.GetGenre(id)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, genre)
}
func CreateGenre(w http.ResponseWriter, r *http.Request) {
	var genre model.Genre

	body, err := helper.ParserBody(r.Body, genre)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.CreateGenre(body.Name)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)

}
func UpdateGenre(w http.ResponseWriter, r *http.Request) {
	var genre model.Genre

	id, err := helper.ParserInt(r, "/genre/update/")
	if err != nil {
		helper.Error(w, err)
	}

	body, err := helper.ParserBody(r.Body, genre)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.UpdateGenre(id, body)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}
func DeleteGenre(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/genre/delete/")
	if err != nil {
		helper.Error(w, err)
	}

	if err := service.DeleteGenre(id); err != nil {
		helper.Error(w, err)
	}
	helper.Success(w, struct {
		Message string `json:"message"`
	}{Message: "genre deleted"})

}
