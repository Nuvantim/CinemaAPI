package service

import(
	db "booking/database"
	model "booking/internal/app/repository"
	ctx "context"
)

func ListBookingSeat(user_id int32)([]model.ListBookingSeatRow,error){

	data,err := db.Queries.ListBookingSeat(ctx.Background(),user_id)
	if err != nil{
		return []model.ListBookingSeatRow{},err
	}
	return data,nil
}

func CreateBookingSeat(body model.CreateBookingSeatParams)([]model.ListBookingSeatRow,error){
	booking_id, err := db.Queries.CreateBookingSeat(ctx.Background(),body)
	if err != nil{
		return []model.ListBookingSeatRow{},err
	}
	data,err := db.Queries.ListBookingSeat(ctx.Background(),booking_id)
	if err != nil{
		return []model.ListBookingSeatRow{},err
	}
	return data,nil
}
