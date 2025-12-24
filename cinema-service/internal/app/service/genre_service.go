package service

import (
	db "cinema/database"
	model "cinema/internal/app/repository"
	ctx "context"
)

func ListGenre() ([]model.Genre, error) {
	genre, err := db.Queries.ListGenre(ctx.Background())
	if err != nil {
		return []model.Genre{}, db.Fatal(err)
	}
	return genre, nil
}

func GetGenre(id int64) (model.Genre, error) {
	genre, err := db.Queries.GetGenre(ctx.Background(), id)
	if err != nil {
		return model.Genre{}, db.Fatal(err)
	}
	return genre, nil
}

func CreateGenre(name string) (model.Genre, error) {
	genre, err := db.Queries.CreateGenre(ctx.Background(), name)
	if err != nil {
		return model.Genre{}, db.Fatal(err)
	}
	return genre, nil
}

func UpdateGenre(body model.UpdateGenreParams) (model.Genre, error) {
	GenreID, err := db.Queries.UpdateGenre(ctx.Background(), body)
	if err != nil {
		return model.Genre{}, db.Fatal(err)
	}

	data, err := GetGenre(GenreID)
	if err != nil {
		return model.Genre{}, db.Fatal(err)
	}

	return data, nil
}

func DeleteGenre(id int64) error {
	if err := db.Queries.DeleteGenre(ctx.Background(), id); err != nil {
		return db.Fatal(err)
	}
	return nil
}
