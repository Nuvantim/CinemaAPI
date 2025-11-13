package service

import (
	db "cinema/database"
	model "cinema/internal/app/repository"
	ctx "context"
)

func ListCinema() ([]model.Cinema, error) {
	data, err := db.Queries.ListCinema(ctx.Background())
	if err != nil {
		return []model.Cinema{}, db.Fatal(err)
	}
	return data, nil
}

func ListCinemaSchedule(id int64) ([]model.ListCinemaScheduleRow, error) {
	data, err := db.Queries.ListCinemaSchedule(ctx.Background(), id)
	if err != nil {
		return []model.ListCinemaScheduleRow{}, db.Fatal(err)
	}
	return data, nil
}

func GetCinema(id int64) (model.Cinema, error) {
	data, err := db.Queries.GetCinema(ctx.Background(), id)
	if err != nil {
		return model.Cinema{}, db.Fatal(err)
	}
	return data, nil
}

func CreateCinema(body model.CreateCinemaParams) (model.Cinema, error) {
	data, err := db.Queries.CreateCinema(ctx.Background(), body)
	if err != nil {
		return model.Cinema{}, db.Fatal(err)
	}
	return data, nil
}

func UpdateCinema(body model.UpdateCinemaParams) (model.Cinema, error) {
	id_cinema, err := db.Queries.UpdateCinema(ctx.Background(), body)
	if err != nil {
		return model.Cinema{}, db.Fatal(err)
	}

	data, err := db.Queries.GetCinema(ctx.Background(), id_cinema)
	if err != nil {
		return model.Cinema{}, db.Fatal(err)
	}

	return data, err

}

func DeleteCinema(id int64) error {
	if err := db.Queries.DeleteCinema(ctx.Background(), id); err != nil {
		return db.Fatal(err)
	}
	return nil
}
