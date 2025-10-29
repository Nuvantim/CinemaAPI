package service

import (
	"fmt"

	model "api/internal/app/request"
	"api/internal/gateway"
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

func CreateGenre(body model.Genre) (model.Genre, error) {
	var endpoint = "/genre/create"

	data, err := gateway.PostCinema[model.Genre, model.Genre](endpoint, body)
	if err != nil {
		return model.Genre{}, err
	}

	return data, nil
}

func UpdateGenre(body model.Genre) (model.Genre, error) {
	var endpoint = "/genre/update"

	data, err := gateway.PutCinema[model.Genre, model.Genre](endpoint, body)
	if err != nil {
		return model.Genre{}, err
	}
	return data, nil
}

func DeleteGenre(id int32) error {
	var endpoint = fmt.Sprintf("/genre/delete/%d", id)

	if err := gateway.DeleteCinema(endpoint); err != nil {
		return err
	}
	return nil
}
