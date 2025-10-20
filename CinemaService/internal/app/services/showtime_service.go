package service

import (
	db "cinema/database"
	model "cinema/internal/app/repository"
	ctx "context"
)

func ListShowTime() ([]model.ListShowTimeRow, error) {
	data, err := db.Queries.ListShowTime(ctx.Background())
	if err != nil {
		return []model.ListShowTimeRow{}, err
	}
	return data, nil
}

func GetShowTime(id int32) (model.GetShowTimeRow, error) {
	data, err := db.Queries.GetShowTime(ctx.Background(), id)
	if err != nil {
		return model.GetShowTimeRow{}, err
	}
	return data, nil
}

func CreateShowTime(body model.CreateShowTimeParams) (model.GetShowTimeRow, error) {
	screen_id, err := db.Queries.CreateShowTime(ctx.Background(), body)
	if err != nil {
		return model.GetShowTimeRow{}, err
	}

	data, err := db.Queries.GetShowTime(ctx.Background(), screen_id)
	if err != nil {
		return model.GetShowTimeRow{}, err
	}

	return data, nil
}

func UpdateShowTime(id int32, body model.UpdateShowTimeParams) (model.GetShowTimeRow, error) {
	body.ID = id

	screen_id, err := db.Queries.UpdateShowTime(ctx.Background(), body)
	if err != nil {
		return model.GetShowTimeRow{}, err
	}

	data, err := db.Queries.GetShowTime(ctx.Background(), screen_id)
	if err != nil {
		return model.GetShowTimeRow{}, err
	}

	return data, nil
}

func DeleteShowTime(id int32) error {
	if err := db.Queries.DeleteShowTime(ctx.Background(), id); err != nil {
		return err
	}

	return nil
}
