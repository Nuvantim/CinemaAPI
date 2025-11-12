package handler

import (
	model "booking/internal/app/repository"
	rds "booking/redis"
	"booking/internal/app/services"
	"booking/pkgs/parser"
	"booking/pkgs/response"

	"net/http"
	"fmt"
)

func ListPayment(w http.ResponseWriter, r *http.Request) {
	var user = struct {
		UserID int64 `json:"user_id"`
	}{}

	body, err := parser.Body(r.Body, user)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.ListPayment(body.UserID)
	if err != nil {
		response.Error(w, err)
		return
	}
	
	defer func(){
	    if data != nil{
	        _ := rds.SetData(fmt.Sprintf("list:payment:%d",user.UserID), data)
	    }
	}()

	response.Success(w, data)
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment model.CreatePaymentParams
	body, err := parser.Body(r.Body, payment)
	if err != nil {
		response.Error(w, err)
		return
	}

	data, err := service.CreatePayment(body)
	if err != nil {
		response.Error(w, err)
		return
	}
	
	
	// set data to redis
    defer func(){
        data_payment,_ := service.ListPayment(data.UserID)
	    if data_payment != nil{
	        _ := rds.SetData(fmt.Sprintf("list:payment:%d",data.UserID), data_payment)
	    }
	}()

	response.Success(w, data)
}