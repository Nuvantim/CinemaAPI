// screen service
package service

import (
	"api/internal/gateway"
	model "cinema/pkg/monorepo"
	"fmt"
)

func ListScreen() ([]model.ListScreenRow, error) {
	url := "/screens"

	data, err := gateway.GetCinema[[]model.ListScreenRow](url)
	if err != nil {
		return []model.ListScreenRow{}, err
	}
	return data, nil
}

func GetScreen(id int64) (model.GetScreenRow, error) {
	url := fmt.Sprintf("/screen/%d", id)

	data, err := gateway.GetCinema[model.GetScreenRow](url)
	if err != nil {
		return model.GetScreenRow{}, err
	}
	return data, nil
}

func CreateScreen(body model.CreateScreenParams) (model.GetScreenRow, error) {
	url := "/screen/create"

	data, err := gateway.PostCinema[model.CreateScreenParams, model.GetScreenRow](url, body)
	if err != nil {
		return model.GetScreenRow{}, err
	}
	return data, nil
}

func UpdateScreen(body model.UpdateScreenParams) (model.GetScreenRow, error) {
	url := "/screen/update"

	data, err := gateway.PutCinema[model.UpdateScreenParams, model.GetScreenRow](url, body)
	if err != nil {
		return model.GetScreenRow{}, err
	}
	return data, nil
}

func DeleteScreen(id int64) error {
	url := fmt.Sprintf("/screen/delete/%d", id)
	if err := gateway.DeleteCinema(url); err != nil {
		return err
	}
	return nil
}
