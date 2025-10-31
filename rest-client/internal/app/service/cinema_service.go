// cinema service
package service

import (
	"api/internal/gateway"
	model "cinema/pkgs/monorepo"
	"fmt"
)

func ListCinema() ([]model.Cinema, error) {
	url := "/cinemas"

	data, err := gateway.GetCinema[[]model.Cinema](url)
	if err != nil {
		return []model.Cinema{}, err
	}
	return data, nil
}

func GetCinema(id int32) (model.Cinema, error) {
	url := fmt.Sprintf("/cinema/%d", id)

	data, err := gateway.GetCinema[model.Cinema](url)
	if err != nil {
		return model.Cinema{}, err
	}

	return data, nil

}

func CreateCinema(body model.CreateCinemaParams) (model.Cinema, error) {
	url := "/cinema/create"

	data, err := gateway.PostCinema[model.CreateCinemaParams, model.Cinema](url, body)
	if err != nil {
		return model.Cinema{}, nil
	}

	return data, nil

}

func UpdateCinema(body model.UpdateCinemaParams) (model.Cinema, error) {
	url := "/cinema/update"

	data, err := gateway.PutCinema[model.UpdateCinemaParams, model.Cinema](url, body)
	if err != nil {
		return model.Cinema{}, nil
	}
	return data, nil

}

func DeleteCinema(id int32) error {
	url := fmt.Sprintf("/cinema/delete/%d", id)

	if err := gateway.DeleteCinema(url); err != nil {
		return err
	}
	return nil

}

