package handlers

import (
	"cinema/internal/app/services"
	"cinema/pkgs/helper"
	"net/http"
)

func ListGenre(w http.ResponseWriter, r *http.Request) {
	genre, err := services.ListGenre()
	if err != nil {
		helper.Error(w, err)
	}
	helper.Success(w, genre)
}
