package service

import (
	db "api/database"
	repo "api/internal/app/repository"
	req "api/internal/app/request"
	"api/pkg/guards"

	ctx "context"
	"errors"
	str "strings"
)

func GetProfile(userID int64) (repo.GetProfileRow, error) {
	data, err := db.Queries.GetProfile(ctx.Background(), userID)
	if err != nil {
		return repo.GetProfileRow{}, errors.New("account not found")
	}
	data.UserAccount.ID = 0
	data.UserProfile.UserID = 0
	data.UserAccount.Password = ""
	return data, nil
}

func UpdateAccount(user req.UpdateAccount, userIDs int64) (repo.GetProfileRow, error) {
	// Define update profile
	var updateAccount = repo.UpdateAccountParams{
		UserID:   userIDs,
		Name:     user.Name,
		Age:      &user.Age,
		Phone:    &user.Phone,
		District: &user.District,
		City:     &user.City,
		Country:  &user.Country,
	}

	// Update password is available
	if str.TrimSpace(user.Password) != "" {
		psw := guard.HashBycrypt(user.Password)
		passUpdate := repo.UpdatePasswordParams{
			ID:       userIDs,
			Password: string(psw),
		}
		if err := db.Queries.UpdatePassword(ctx.Background(), passUpdate); err != nil {
			return repo.GetProfileRow{}, db.Fatal(err)
		}
	}
	// Update Profile User
	if err := db.Queries.UpdateAccount(ctx.Background(), updateAccount); err != nil {
		return repo.GetProfileRow{}, db.Fatal(err)
	}

	// Returning data
	usr, err := GetProfile(userIDs)
	if err != nil {
		return repo.GetProfileRow{}, db.Fatal(err)
	}
	return usr, nil
}

func DeleteAccount(userID int64) (string, error) {
	if err := db.Queries.DeleteAccount(ctx.Background(), userID); err != nil {
		return "", db.Fatal(err)
	}
	return "Your account successfuly delete", nil
}
