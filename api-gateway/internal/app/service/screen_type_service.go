package service

import (
	"api/internal/gateway"
	model "cinema/pkg/monorepo"
	"fmt"
)

func ListScreenType() ([]model.ScreenType, error) {
	url := "/screen/types"

	data, err := gateway.GetCinema[[]model.ScreenType](url)
	if err != nil {
		return []model.ScreenType{}, err
	}
	return data, nil
}

func GetScreenType(id int64) (model.ScreenType, error) {
	url := fmt.Sprintf("/screen/type/%d", id)

	data, err := gateway.GetCinema[model.ScreenType](url)
	if err != nil {
		return model.ScreenType{}, err
	}
	return data, nil
}

func CreateScreenType(body any) (model.ScreenType, error) {
	url := "/screen/type/create"

	data, err := gateway.PostCinema[any, model.ScreenType](url, body)
	if err != nil {
		return model.ScreenType{}, err
	}
	return data, nil
}

func UpdateScreenType(body model.ScreenType) (model.ScreenType, error) {
	url := "/screen/type/update"

	data, err := gateway.PutCinema[model.ScreenType, model.ScreenType](url, body)
	if err != nil {
		return model.ScreenType{}, err
	}
	return data, nil
}

func DeleteScreenType(id int64) error {
	url := fmt.Sprintf("/screen/type/delete/%d", id)

	if err := gateway.DeleteCinema(url); err != nil {
		return err
	}
	return nil

}
