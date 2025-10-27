package service

import (
	"api/internal/gateway"
	model "github.com/Nuvantim/CinemaService/internal/app/repository"
)

func ListGenre() ([]model.Genre, error) {
	var genre []model.Genre
	var endpoint = "/genres"

	data, err := gateway.GetCinema(endpoint, genre)
	if err != nil {
		return []model.Genre{}, err
	}
	return data, nil
}

func GetGenre(id int32) (model.Genre, error) {
	var genre model.Genre
	var endpoint = fmt.Sprintf("/genre/%d", id)

	data, err := gateway.GetCinema(endpoint, genre)
	if err != nil {
		return model.Genre{}, err
	}

	return data, nil
}