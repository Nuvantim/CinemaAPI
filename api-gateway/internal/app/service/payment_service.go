// payment service
package service

import (
	"api/internal/gateway"
	model "booking/pkgs/monorepo"
 	rds "api/redis"

 	"fmt"
)

func ListPayment(body any) ([]model.Payment, error) {
	// check data on redis
	value := fmt.Sprintf("list:booking:%d", body.UserID)
	redis_data, err := rds.GetData[*[]model.Payment](value)
	if err == nil && redis_data != nil {
		return redis_data,nil
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
