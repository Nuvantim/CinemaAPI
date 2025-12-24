package handler

import (
	model "cinema/internal/app/repository"
	"cinema/internal/app/service"
	"cinema/pkg/parser"
	"cinema/pkg/response"
	rds "cinema/redis"
	"net/http"
)

func ListShowTime(w http.ResponseWriter, r *http.Request) {
	data, err := service.ListShowTime()
	if err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		if data != nil {
			_ = rds.SetData("list:showtime", data)
		}
	}()

	response.Success(w, data)
}

func GetShowTime(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/showtime/")
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.GetShowTime(id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, data)
}

func CreateShowTime(w http.ResponseWriter, r *http.Request) {
	var showtime model.CreateShowTimeParams
	body, err := parser.Body(r.Body, showtime)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.CreateShowTime(body)
	if err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		data_showtime, _ := service.ListShowTime()
		if data_showtime != nil {
			_ = rds.SetData("list:showtime", data_showtime)
		}
	}()

	response.Success(w, data)
}

func UpdateShowTime(w http.ResponseWriter, r *http.Request) {
	var showtime model.UpdateShowTimeParams
	body, err := parser.Body(r.Body, showtime)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.UpdateShowTime(body)
	if err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		data_showtime, _ := service.ListShowTime()
		if data_showtime != nil {
			_ = rds.SetData("list:showtime", data_showtime)
		}
	}()

	response.Success(w, data)

}

func DeleteShowTime(w http.ResponseWriter, r *http.Request) {
	id, err := parser.ParamsInt(r, "/showtime/delete/")
	if err != nil {
		response.Error(w, err)
		return
	}

	if err := service.DeleteShowTime(id); err != nil {
		response.Error(w, err)
		return
	}

	// set data to redis
	go func() {
		data_showtime, _ := service.ListShowTime()
		if data_showtime != nil {
			_ = rds.SetData("list:showtime", data_showtime)
		}
	}()

	response.Success(w, struct {
		Message string `json:"message"`
	}{Message: "showtime deleted"})
}
