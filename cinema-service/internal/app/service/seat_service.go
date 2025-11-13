package service

import (
	db "cinema/database"
	model "cinema/internal/app/repository"
	ctx "context"
)

func ListSeat() ([]model.ListSeatRow, error) {
	data, err := db.Queries.ListSeat(ctx.Background())
	if err != nil {
		return []model.ListSeatRow{}, db.Fatal(err)
	}
	return data, nil
}

func GetSeat(id int64) (model.GetSeatRow, error) {
	data, err := db.Queries.GetSeat(ctx.Background(), id)
	if err != nil {
		return model.GetSeatRow{}, db.Fatal(err)
	}

	return data, nil
}

func CreateSeat(body model.CreateSeatParams) (model.GetSeatRow, error) {
	seat_id, err := db.Queries.CreateSeat(ctx.Background(), body)
	if err != nil {
		return model.GetSeatRow{}, db.Fatal(err)
	}

	data, err := db.Queries.GetSeat(ctx.Background(), seat_id)
	if err != nil {
		return model.GetSeatRow{}, db.Fatal(err)
	}

	return data, nil
}

func SeatPrice(body model.SeatPriceParams) (int64, error) {
	data, err := db.Queries.SeatPrice(ctx.Background(), body)
	if err != nil {
		return 0, db.Fatal(err)
	}

	return data, nil
}

func UpdateSeat(body model.UpdateSeatParams) (model.GetSeatRow, error) {
	seat_id, err := db.Queries.UpdateSeat(ctx.Background(), body)
	if err != nil {
		return model.GetSeatRow{}, db.Fatal(err)
	}

	data, err := db.Queries.GetSeat(ctx.Background(), seat_id)
	if err != nil {
		return model.GetSeatRow{}, db.Fatal(err)
	}

	return data, nil

}

func DeleteSeat(id int64) error {
	if err := db.Queries.DeleteSeat(ctx.Background(), id); err != nil {
		return db.Fatal(err)
	}

	return nil
}
