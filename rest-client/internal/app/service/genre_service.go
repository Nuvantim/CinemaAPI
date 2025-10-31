package service

import (
	"fmt"

	"api/internal/gateway"
	model "cinema/pkgs/monorepo"
)

func ListGenre() ([]model.Genre, error) {
	url := "/genres"

	data, err := gateway.GetCinema[[]model.Genre](url)
	if err != nil {
		return []model.Genre{}, err
	}
	return data, nil
}

func GetGenre(id int32) (model.Genre, error) {
	url := fmt.Sprintf("/genre/%d", id)

	data, err := gateway.GetCinema[model.Genre](url)
	if err != nil {
		return model.Genre{}, err
	}

	return data, nil
}

func CreateGenre(body model.Genre) (model.Genre, error) {
	var url = "/genre/create"

	data, err := gateway.PostCinema[model.Genre, model.Genre](url, body)
	if err != nil {
		return model.Genre{}, err
	}

	return data, nil
}

func UpdateGenre(body model.Genre) (model.Genre, error) {
	var url = "/genre/update"

	data, err := gateway.PutCinema[model.Genre, model.Genre](url, body)
	if err != nil {
		return model.Genre{}, err
	}
	return data, nil
}

func DeleteGenre(id int32) error {
	var url = fmt.Sprintf("/genre/delete/%d", id)

	if err := gateway.DeleteCinema(url); err != nil {
		return err
	}
	return nil
}
