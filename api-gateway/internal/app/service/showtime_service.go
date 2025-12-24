// showtime service
package service

import (
	"fmt"

	"api/internal/gateway"
	model "cinema/pkg/monorepo"
)

func ListShowTime() ([]model.ListShowTimeRow, error) {
	url := "/showtimes"

	data, err := gateway.GetCinema[[]model.ListShowTimeRow](url)
	if err != nil {
		return []model.ListShowTimeRow{}, err
	}
	return data, nil
}

func GetShowTime(id int64) (model.GetShowTimeRow, error) {
	url := fmt.Sprintf("/showtime/%d", id)

	data, err := gateway.GetCinema[model.GetShowTimeRow](url)
	if err != nil {
		return model.GetShowTimeRow{}, err
	}
	return data, nil

}

func CreateShowTime(body model.CreateShowTimeParams) (model.GetShowTimeRow, error) {
	url := "/showtime/create"

	data, err := gateway.PostCinema[model.CreateShowTimeParams, model.GetShowTimeRow](url, body)
	if err != nil {
		return model.GetShowTimeRow{}, err
	}
	return data, nil
}

func UpdateShowTime(body model.UpdateShowTimeParams) (model.GetShowTimeRow, error) {
	url := "/showtime/update"

	data, err := gateway.PutCinema[model.UpdateShowTimeParams, model.GetShowTimeRow](url, body)
	if err != nil {
		return model.GetShowTimeRow{}, err
	}

	return data, nil
}

func DeleteShowTime(id int64) error {
	url := fmt.Sprintf("/showtime/delete/%d", id)

	if err := gateway.DeleteCinema(url); err != nil {
		return err
	}
	return nil

}
