package services

import (
	db "cinema/database"
	model "cinema/internal/app/repository"
	ctx "context"
)

func ListGenre() ([]model.Genre, error) {
	genre, err := db.Queries.ListGenre(ctx.Background())
	if err != nil {
		return []model.Genre{}, err
	}
	return genre, nil
}
