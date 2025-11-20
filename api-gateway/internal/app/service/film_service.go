package service

import (
	"api/internal/gateway"
	model "cinema/pkg/monorepo"
	"fmt"
)

func ListFilm() ([]model.ListFilmRow, error) {
	url := "/films"

	data, err := gateway.GetCinema[[]model.ListFilmRow](url)
	if err != nil {
		return []model.ListFilmRow{}, err
	}
	return data, nil
}

func GetFilm(id int64) (model.GetFilmRow, error) {
	url := fmt.Sprintf("/film/%d", id)

	data, err := gateway.GetCinema[model.GetFilmRow](url)
	if err != nil {
		return model.GetFilmRow{}, err
	}
	return data, nil
}

func SearchFilm(body any) ([]model.SearchFilmRow, error) {
	url := "/film/search"
	data, err := gateway.PostCinema[any, []model.SearchFilmRow](url, body)
	if err != nil {
		return []model.SearchFilmRow{}, err
	}
	return data, nil
}

func SearchFilmGenre(id int64) ([]model.SearchGenreFilmRow, error) {
	url := fmt.Sprintf("/film/genre/%d", id)
	data, err := gateway.GetCinema[[]model.SearchGenreFilmRow](url)
	if err != nil {
		return []model.SearchGenreFilmRow{}, err
	}
	return data, nil
}

func CreateFilm(body model.CreateFilmParams) (model.GetFilmRow, error) {
	url := "/film/create"
	data, err := gateway.PostCinema[model.CreateFilmParams, model.GetFilmRow](url, body)
	if err != nil {
		return model.GetFilmRow{}, err
	}
	return data, nil
}

func UpdateFilm(body model.UpdateFilmParams) (model.GetFilmRow, error) {
	url := "/film/update"
	data, err := gateway.PutCinema[model.UpdateFilmParams, model.GetFilmRow](url, body)
	if err != nil {
		return model.GetFilmRow{}, err
	}
	return data, nil
}

func DeleteFilm(id int64) error {
	url := fmt.Sprintf("/film/delete/%d", id)
	if err := gateway.DeleteCinema(url); err != nil {
		return err
	}
	return nil
}
