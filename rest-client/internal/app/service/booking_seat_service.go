// booking seat service
package service

import (
	"fmt"

	"api/internal/gateway"
	model "booking/pkgs/monorepo"
)

func ListBookingSeat(body any) ([]model.BookingSeat, error) {
	url := "/booking/seats"

	data, err := gateway.PostBooking[any, []model.BookingSeat](url, body)
	if err != nil {
		return []model.BookingSeat{}, err
	}

	return data, nil
}

func CreateBookingSeat(body model.CreateBookingSeatParams) (model.BookingSeat, error) {
	url := "/booking/seat/create"

	data, err := gateway.PostBooking[model.CreateBookingSeatParams, model.BookingSeat](url, body)
	if err != nil {
		return model.BookingSeat{}, err
	}

	return data, nil
}

func DeleteBookingSeat(id int64) error {
	url := fmt.Sprintf("/booking/seat/delete/%d", id)

	if err := gateway.DeleteBookingSeat(url); err != nil {
		return err
	}
	return nil
}
