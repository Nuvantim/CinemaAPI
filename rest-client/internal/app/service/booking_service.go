// booking service
package service

import (
	"fmt"

	"api/internal/gateway"
	model "booking/pkgs/monorepo"
)

func ListBooking(body any) ([]model.Booking, error) {
	url := "/bookings"

	data, err := gateway.PostBooking[any, []model.Booking](url, body)
	if err != nil {
		return []model.Booking{}, err
	}

	return data, nil
}

func CreateBooking(body model.CreateBookingParams) (model.Booking, error) {
	var url = "/booking/create"

	data, err := gateway.PostBooking[model.CreateBookingParams, model.Booking](url, body)
	if err != nil {
		return model.Booking{}, err
	}

	return data, nil
}

func DeleteBooking(id int64) error {
	var url = fmt.Sprintf("/booking/delete/%d", id)

	if err := gateway.DeleteBooking(url); err != nil {
		return err
	}
	return nil
}
