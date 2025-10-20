package service

import (
	db "cinema/database"
	model "cinema/internal/app/repository"
	ctx "context"
)

func ListSeat() ([]model.ListSeatRow, error) {
	data, err := db.Queries.ListSeat(ctx.Background())
	if err != nil {
		return []model.ListSeatRow{}, err
	}
	return data, nil
}

func GetSeat(id int32) (model.GetSeatRow, error) {
	data, err := db.Queries.GetSeat(ctx.Background(), id)
	if err != nil {
		return model.GetSeatRow{}, err
	}

	return data, nil
}

func CreateSeat(body model.CreateSeatParams) (model.GetSeatRow, error) {
	seat_id, err := db.Queries.CreateSeat(ctx.Background(), body)
	if err != nil {
		return model.GetSeatRow{}, err
	}

	data, err := db.Queries.GetSeat(ctx.Background(), seat_id)
	if err != nil {
		return model.GetSeatRow{}, err
	}

	return data, nil
}

func UpdateSeat(id int32, body model.UpdateSeatParams) (model.GetSeatRow, error) {
	body.ID = id

	seat_id, err := db.Queries.UpdateSeat(ctx.Background(), body)
	if err != nil {
		return model.GetSeatRow{}, err
	}

	data, err := db.Queries.GetSeat(ctx.Background(), seat_id)
	if err != nil {
		return model.GetSeatRow{}, err
	}

	return data, nil

}

func DeleteSeat(id int32) error {
	if err := db.Queries.DeleteSeat(ctx.Background(), id); err != nil {
		return err
	}

	return nil
}
