package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/services"
	"cinema/pkgs/helper"
	"net/http"
)

func ListFilm(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListFilm()
	if err != nil {
		helper.Error(w, err)
	}
	helper.Success(w, data)
}
func GetFilm(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/film/")
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.GetFilm(id)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}
func SearchFilm(w http.ResponseWriter, r *http.Request) {
	var film model.Film
	body, err := helper.ParserBody(r.Body, film)
	if err != nil {
		helper.Error(w, err)
	}
	data, err := service.SearchFilm(body.Title)
	if err != nil {
		helper.Error(w, err)
	}
	helper.Success(w, data)
}
func SearchFilmGenre(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/film/genre/")
	if err != nil {
		helper.Error(w, err)
	}
	data, err := service.SearchFilmGenre(id)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}
func CreateFilm(w http.ResponseWriter, r *http.Request) {
	var film model.CreateFilmParams

	body, err := helper.ParserBody(r.Body, film)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.CreateFilm(body)
	if err != nil {
		helper.Error(w, err)
	}
	helper.Success(w, data)

}
func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/film/update/")
	if err != nil {
		helper.Error(w, err)
	}

	var film model.UpdateFilmParams

	body, err := helper.ParserBody(r.Body, film)
	if err != nil {
		helper.Error(w, err)
	}

	data, err := service.UpdateFilm(id, body)
	if err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, data)
}
func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParserInt(r, "/film/delete/")
	if err != nil {
		helper.Error(w, err)
	}

	if err := service.DeleteFilm(id); err != nil {
		helper.Error(w, err)
	}

	helper.Success(w, struct {
		Message string `json:"message"`
	}{Message: "film deleted"})

}
