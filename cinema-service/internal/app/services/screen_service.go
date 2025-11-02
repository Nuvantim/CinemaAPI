package service

import (
	db "cinema/database"
	model "cinema/internal/app/repository"
	ctx "context"
)

func ListScreen() ([]model.ListScreenRow, error) {
	data, err := db.Queries.ListScreen(ctx.Background())
	if err != nil {
		return []model.ListScreenRow{}, db.Fatal(err)
	}
	return data, nil
}

func GetScreen(id int64) (model.GetScreenRow, error) {
	data, err := db.Queries.GetScreen(ctx.Background(), id)
	if err != nil {
		return model.GetScreenRow{}, db.Fatal(err)
	}
	return data, nil
}

func CreateScreen(body model.CreateScreenParams) (model.GetScreenRow, error) {
	screen_id, err := db.Queries.CreateScreen(ctx.Background(), body)
	if err != nil {
		return model.GetScreenRow{}, db.Fatal(err)
	}

	data, err := db.Queries.GetScreen(ctx.Background(), screen_id)
	if err != nil {
		return model.GetScreenRow{}, db.Fatal(err)
	}

	return data, nil
}

func UpdateScreen(body model.UpdateScreenParams) (model.GetScreenRow, error) {

	screen_id, err := db.Queries.UpdateScreen(ctx.Background(), body)
	if err != nil {
		return model.GetScreenRow{}, db.Fatal(err)
	}

	data, err := db.Queries.GetScreen(ctx.Background(), screen_id)
	if err != nil {
		return model.GetScreenRow{}, db.Fatal(err)
	}

	return data, nil
}

func DeleteScreen(id int64) error {
	if err := db.Queries.DeleteScreen(ctx.Background(), id); err != nil {
		return db.Fatal(err)
	}

	return nil
}
