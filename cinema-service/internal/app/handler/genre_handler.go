package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/service"
	"cinema/pkg/parser"
	"cinema/pkg/response"
	"net/http"
)

func ListGenre(w http.ResponseWriter, r *http.Request) {
	genre, err := service.ListGenre()
	if err != nil {
		response.Error(w, err)
	}
	response.Success(w, genre)
}

func GetGenre(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/genre/")
	if err != nil {
		response.Error(w, err)
		return
	}

	genre, err := service.GetGenre(id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, genre)
}
func CreateGenre(w http.ResponseWriter, r *http.Request) {
	var genre model.Genre

	body, err := parser.Body(r.Body, genre)
	if err != nil {
		response.Error(w, err)
	}

	data, err := service.CreateGenre(body.Name)
	if err != nil {
		response.Error(w, err)
	}

	response.Success(w, data)

}
func UpdateGenre(w http.ResponseWriter, r *http.Request) {
	var genre model.UpdateGenreParams

	body, err := parser.Body(r.Body, genre)
	if err != nil {
		response.Error(w, err)
	}

	data, err := service.UpdateGenre(body)
	if err != nil {
		response.Error(w, err)
	}

	response.Success(w, data)
}
func DeleteGenre(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/genre/delete/")
	if err != nil {
		response.Error(w, err)
	}

	if err := service.DeleteGenre(id); err != nil {
		response.Error(w, err)
	}
	response.Success(w, struct {
		Message string `json:"message"`
	}{Message: "genre deleted"})

}
