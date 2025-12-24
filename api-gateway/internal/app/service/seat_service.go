// seat service
package service

import (
	"fmt"

	"api/internal/gateway"
	model "cinema/pkg/monorepo"
)

func ListSeat() ([]model.ListSeatRow, error) {
	url := "/seats"

	data, err := gateway.GetCinema[[]model.ListSeatRow](url)
	if err != nil {
		return []model.ListSeatRow{}, err
	}
	return data, nil
}

func GetSeat(id int64) (model.GetSeatRow, error) {
	url := fmt.Sprintf("/seat/%d", id)

	data, err := gateway.GetCinema[model.GetSeatRow](url)
	if err != nil {
		return model.GetSeatRow{}, err
	}
	return data, nil
}

func CreateSeat(body model.CreateSeatParams) (model.GetSeatRow, error) {
	url := "/seat/create"

	data, err := gateway.PostCinema[model.CreateSeatParams, model.GetSeatRow](url, body)
	if err != nil {
		return model.GetSeatRow{}, err
	}
	return data, nil
}

func UpdateSeat(body model.UpdateSeatParams) (model.GetSeatRow, error) {
	url := "/seat/update"

	data, err := gateway.PutCinema[model.UpdateSeatParams, model.GetSeatRow](url, body)
	if err != nil {
		return model.GetSeatRow{}, err
	}
	return data, nil
}

func SeatPrice(showtime_id, seat_id int64) (int64, error) {
	var seat = model.SeatPriceParams{
		ShowtimeID: showtime_id,
		SeatID:     seat_id,
	}

	url := "/seat/price"
	data, err := gateway.PostCinema[model.SeatPriceParams, int64](url, seat)
	if err != nil {
		return 0, err
	}

	return data, nil
}

func DeleteSeat(id int64) error {
	url := fmt.Sprintf("/seat/delete/%d", id)

	if err := gateway.DeleteCinema(url); err != nil {
		return err
	}
	return nil
}
