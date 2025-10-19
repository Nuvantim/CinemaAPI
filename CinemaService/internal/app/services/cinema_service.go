package service

import (
	db "cinema/database"
	model "cinema/internal/app/repository"
	ctx "context"
)

func ListCinema() ([]model.Cinema, error) {
	data, err := db.Queries.ListCinema(ctx.Background())
	if err != nil {
		return []model.Cinema{}, err
	}
	return data, nil
}

func GetCinema(id int32) (model.Cinema, error) {
	data, err := db.Queries.GetCinema(ctx.Background(), id)
	if err != nil {
		return model.Cinema{}, err
	}
	return data, nil
}

func CreateCinema(body model.CreateCinemaParams) (model.Cinema, error) {
	data, err := db.Queries.CreateCinema(ctx.Background(), body)
	if err != nil {
		return model.Cinema{}, err
	}
	return data, nil
}

func UpdateCinema(id int32, body model.UpdateCinemaParams) (model.Cinema, error) {
	body.ID = id

	id_cinema, err := db.Queries.UpdateCinema(ctx.Background(), body)
	if err != nil {
		return model.Cinema{}, err
	}

	data, err := db.Queries.GetCinema(ctx.Background(), id_cinema)
	if err != nil {
		return model.Cinema{}, err
	}

	return data, err

}

func DeleteCinema(id int32) error {
	if err := db.Queries.DeleteCinema(ctx.Background(), id); err != nil {
		return err
	}
	return nil
}
