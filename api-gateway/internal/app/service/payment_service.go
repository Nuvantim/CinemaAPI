// payment service
package service

import (
	"api/internal/gateway"
	model "booking/pkgs/monorepo"
)

func ListPayment(body any) ([]model.Payment, error) {
	url := "/payments"

	data, err := gateway.PostBooking[any, []model.Payment](url, body)
	if err != nil {
		return []model.Payment{}, err
	}

	return data, nil
}

func CreatePayment(body model.CreatePaymentParams) (model.Payment, error) {
	url := "/payment/create"

	data, err := gateway.PostBooking[model.CreatePaymentParams, model.Payment](url, body)
	if err != nil {
		return model.Payment{}, err
	}

	return data, nil
}
