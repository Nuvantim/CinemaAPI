package service

import(
	db "cinema/database"
	model "cinema/internal/app/repository"
	ctx "context"
)

func ListScreen()([]ListScreenRow,error){}

func GetScreen(id int)(model.GetScreenRow,error){}

func CreateScreen(body model.CreateScreenParams)(model.GetScreenRow,error){}

func UpdateScreen(id int32, model.UpdateScreenParams)(model.GetScreenRow,error){}

func DeleteScreen(id int32)error{}

