package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/service"
	"cinema/pkg/parser"
	"cinema/pkg/response"
	rds "cinema/redis"

	"net/http"
)

func ListCinema(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListCinema()
	if err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		if data != nil {
			_ = rds.SetData("list:cinema", data)
		}
	}()

	response.Success(w, data)
}

func ListCinemaSchedule(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/cinema/schedules/")
	if err != nil {
		response.Error(w, err)
		return
	}
	data, err := service.ListCinemaSchedule(id)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.Success(w, data)
}

func GetCinema(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/cinema/")
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.GetCinema(id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func CreateCinema(w http.ResponseWriter, r *http.Request) {
	var cinema model.CreateCinemaParams
	body, err := parser.Body(r.Body, cinema)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.CreateCinema(body)
	if err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		data_cinema, _ := service.ListCinema()
		if data_cinema != nil {
			_ = rds.SetData("list:cinema", data)
		}
	}()

	response.Success(w, data)
}

func UpdateCinema(w http.ResponseWriter, r *http.Request) {
	var cinema model.UpdateCinemaParams

	body, err := parser.Body(r.Body, cinema)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.UpdateCinema(body)
	if err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		data_cinema, _ := service.ListCinema()
		if data_cinema != nil {
			_ = rds.SetData("list:cinema", data)
		}
	}()

	response.Success(w, data)
}

func DeleteCinema(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/cinema/delete/")
	if err != nil {
		response.Error(w, err)
		return
	}

	if err := service.DeleteCinema(id); err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		data_cinema, _ := service.ListCinema()
		if data_cinema != nil {
			_ = rds.SetData("list:cinema", data_cinema)
		}
	}()

	response.Success(w, struct {
		Message string `json:"message"`
	}{Message: "cinema deleted"})
}
