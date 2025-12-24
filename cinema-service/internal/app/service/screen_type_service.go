package service

import (
	db "cinema/database"
	model "cinema/internal/app/repository"
	ctx "context"
)

func ListScreenType() ([]model.ScreenType, error) {
	data, err := db.Queries.ListScreenType(ctx.Background())
	if err != nil {
		return []model.ScreenType{}, db.Fatal(err)
	}

	return data, nil
}

func GetScreenType(id int64) (model.ScreenType, error) {
	data, err := db.Queries.GetScreenType(ctx.Background(), id)
	if err != nil {
		return model.ScreenType{}, db.Fatal(err)
	}

	return data, err
}

func CreateScreenType(name string) (model.ScreenType, error) {
	screen_type_id, err := db.Queries.CreateScreenType(ctx.Background(), name)
	if err != nil {
		return model.ScreenType{}, db.Fatal(err)
	}

	data, err := db.Queries.GetScreenType(ctx.Background(), screen_type_id)
	if err != nil {
		return model.ScreenType{}, db.Fatal(err)
	}

	return data, nil
}

func UpdateScreenType(body model.UpdateScreenTypeParams) (model.ScreenType, error) {
	screen_type_id, err := db.Queries.UpdateScreenType(ctx.Background(), body)
	if err != nil {
		return model.ScreenType{}, db.Fatal(err)
	}

	data, err := db.Queries.GetScreenType(ctx.Background(), screen_type_id)
	if err != nil {
		return model.ScreenType{}, db.Fatal(err)
	}

	return data, nil
}

func DeleteScreenType(id int64) error {
	if err := db.Queries.DeleteScreenType(ctx.Background(), id); err != nil {
		return db.Fatal(err)
	}
	return nil
}
