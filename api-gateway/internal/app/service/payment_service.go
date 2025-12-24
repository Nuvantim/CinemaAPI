// payment service
package service

import (
	req "api/internal/app/request"
	"api/internal/gateway"
	rds "api/redis"
	model "booking/pkg/monorepo"

	"fmt"
)

func ListPayment(body req.BookingPayment) ([]model.Payment, error) {
	// check data on redis
	key := fmt.Sprintf("list:payment:%d", body.UserID)
	redis_data, err := rds.GetData[[]model.Payment](key)
	if err == nil && redis_data != nil {
		return *redis_data, nil
	}

	// get data from service
	url := "/payment/"

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
