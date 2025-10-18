package service

import (
	db "cinema/database"
	model "cinema/internal/app/repository"
	ctx "context"
)

func ListFilm() ([]model.ListFilmRow, error) {
	data, err := db.Queries.ListFilm(ctx.Background())
	if err != nil {
		return model.ListFilmRow{}, err
	}
	return data, nil
}
func GetFilm(id int32) (model.GetFilmRow, error) {
	data, err := db.Queries.GetFilm(ctx.Background(), id)
	if err != nil {
		return model.GetFilmRow{}, err
	}

	return data, nil
}
func SearchFilm(title string) ([]model.SearchFilmRow, error) {
	data, err := db.Queries.SearchFilm(ctx.Background(), title)
	if err != nil {
		return model.SearchFilmRow{}, err
	}
	return data, nil
}
func SearchFilmGenre(id int32) ([]model.SearchGenreFilmRow, error) {
	data, err := db.Queries.SearchGenreFilm(ctx.Background(), id)
	if err != nil {
		return model.SearchGenreFilmRow{}, err
	}
	return data, nil
}

// func CreateFilm()(,error){}
// func UpdateFilm()(,error){}
// func DeleteFilm()(,error){}
