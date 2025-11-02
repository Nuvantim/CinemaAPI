package service

import (
	db "cinema/database"
	model "cinema/internal/app/repository"
	ctx "context"
)

func ListFilm() ([]model.ListFilmRow, error) {
	data, err := db.Queries.ListFilm(ctx.Background())
	if err != nil {
		return []model.ListFilmRow{}, db.Fatal(err)
	}
	return data, nil
}

func GetFilm(id int64) (model.GetFilmRow, error) {
	data, err := db.Queries.GetFilm(ctx.Background(), id)
	if err != nil {
		return model.GetFilmRow{}, db.Fatal(err)
	}

	return data, nil
}

func SearchFilm(title string) ([]model.SearchFilmRow, error) {
	data, err := db.Queries.SearchFilm(ctx.Background(), title)
	if err != nil {
		return []model.SearchFilmRow{}, db.Fatal(err)
	}
	return data, nil
}

func SearchFilmGenre(id int64) ([]model.SearchGenreFilmRow, error) {
	data, err := db.Queries.SearchGenreFilm(ctx.Background(), id)
	if err != nil {
		return []model.SearchGenreFilmRow{}, db.Fatal(err)
	}
	return data, nil
}

func CreateFilm(body model.CreateFilmParams) (model.GetFilmRow, error) {
	id_film, err := db.Queries.CreateFilm(ctx.Background(), body)
	if err != nil {
		return model.GetFilmRow{}, db.Fatal(err)
	}
	data, err := db.Queries.GetFilm(ctx.Background(), id_film)
	if err != nil {
		return model.GetFilmRow{}, db.Fatal(err)
	}

	return data, nil
}

func UpdateFilm(body model.UpdateFilmParams) (model.GetFilmRow, error) {
	id_film, err := db.Queries.UpdateFilm(ctx.Background(), body)
	if err != nil {
		return model.GetFilmRow{}, db.Fatal(err)
	}

	data, err := db.Queries.GetFilm(ctx.Background(), id_film)
	if err != nil {
		return model.GetFilmRow{}, db.Fatal(err)
	}

	return data, nil

}

func DeleteFilm(id int64) error {
	if err := db.Queries.DeleteFilm(ctx.Background(), id); err != nil {
		return db.Fatal(err)
	}
	return nil
}
