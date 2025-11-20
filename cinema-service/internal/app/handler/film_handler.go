package handler

import (
	model "cinema/internal/app/repository"
	rds "cinema/redis"

	"cinema/internal/app/service"
	"cinema/pkg/parser"
	"cinema/pkg/response"
	"cinema/pkg/validator"

	"net/http"
)

func ListFilm(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListFilm()
	if err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		if data != nil {
			_ = rds.SetData("list:film", data)
		}
	}()

	response.Success(w, data)
}

func GetFilm(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/film/")
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.GetFilm(id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func SearchFilm(w http.ResponseWriter, r *http.Request) {
	var film = struct {
		Title string `json:"title"`
	}{}

	body, err := parser.Body(r.Body, film)
	if err != nil {
		response.Error(w, err)
		return
	}
	data, err := service.SearchFilm(body.Title)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Success(w, data)
}

func SearchFilmGenre(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/film/genre/")
	if err != nil {
		response.Error(w, err)
		return
	}
	data, err := service.SearchFilmGenre(id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	var film model.CreateFilmParams

	body, err := parser.Body(r.Body, film)
	if err != nil {
		response.Error(w, err)
		return
	}

	// duration validation
	if err := validator.DurationValid(body.Duration); err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.CreateFilm(body)
	if err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		data_film, _ := service.ListFilm()
		if data_film != nil {
			_ = rds.SetData("list:film", data_film)
		}
	}()

	response.Success(w, data)

}

func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	var film model.UpdateFilmParams

	body, err := parser.Body(r.Body, film)
	if err != nil {
		response.Error(w, err)
		return
	}

	// duration validation
	if err := validator.DurationValid(body.Duration); err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.UpdateFilm(body)
	if err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		data_film, _ := service.ListFilm()
		if data_film != nil {
			_ = rds.SetData("list:film", data_film)
		}
	}()

	response.Success(w, data)
}

func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/film/delete/")
	if err != nil {
		response.Error(w, err)
		return
	}

	if err := service.DeleteFilm(id); err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		data_film, _ := service.ListFilm()
		if data_film != nil {
			_ = rds.SetData("list:film", data_film)
		}
	}()

	response.Success(w, struct {
		Message string `json:"message"`
	}{Message: "film deleted"})

}
