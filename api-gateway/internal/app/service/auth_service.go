package service

import (
	db "api/database"
	repo "api/internal/app/repository"
	req "api/internal/app/request"
	"api/pkg/guards"
	rds "api/redis"

	ctx "context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func SendOTP(email string) (string, error) {
	// generate otp
	otp := guard.GenerateOTP()

	// create struct data
	token := req.CreateOTP{
		Code:  otp,
		Email: email,
	}

	// set data to redis
	var key string = fmt.Sprintf("verif:%s", token.Code)
	if err := rds.SetData(key, token); err != nil {
		return "", err
	}

	// send otp via email
	if error := guard.SendOTP(token.Email, token.Code); error != nil {
		return "", error
	}

	return "otp send successfully", nil

}

func Register(regist req.Register) (string, error) {
	// search otp
	var key string = fmt.Sprintf("verif:%s", regist.Code)
	result, err := rds.GetData[req.CreateOTP](key)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", fmt.Errorf("OTP not found or expired")
	}

	// check email
	if result.Email != regist.Email {
		return "", fmt.Errorf("registration rejected")
	}

	_, err = db.Queries.FindEmail(ctx.Background(), result.Email)
	if err == nil {
		return "", errors.New("email already exist")
	}

	var pass = guard.HashBycrypt(regist.Password) // Hashing Password
	// Regist New User
	createUser := repo.CreateUserParams{
		Name:     regist.Name,
		Email:    regist.Email,
		Password: string(pass),
	}

	// create user_account
	if err := db.Queries.CreateUser(ctx.Background(), createUser); err != nil {
		return "", db.Fatal(err)
	}

	// Delete the OTP on redis
	_ = rds.DelData(key)

	return "your account has been created, please login", nil
}

func Login(login req.Login) (string, string, error) {
	// Find data account
	data, err := db.Queries.FindEmail(ctx.Background(), login.Email)
	if err != nil {
		return "", "", errors.New("account not found")
	}
	//compared password
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(login.Password)); err != nil {
		return "", "", errors.New("password incorect")
	}

	role, err := db.Queries.AllRoleClient(ctx.Background(), data.ID)
	if err != nil {
		return "", "", db.Fatal(err)
	}

	// Create access token and refresh token
	accessToken, err := guard.CreateToken(data.ID, data.Email, role)
	if err != nil {
		return "", "", db.Fatal(err)
	}

	refreshToken, err := guard.CreateRefreshToken(data.ID, data.Email)
	if err != nil {
		return "", "", db.Fatal(err)
	}

	return accessToken, refreshToken, nil

}

func ResetPassword(pass req.ResetPassword) (string, error) {
	// Check Code Otp
	var key string = fmt.Sprintf("verif:%s", pass.Code)
	result, err := rds.GetData[req.CreateOTP](key)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", fmt.Errorf("OTP not found or expired")
	}

	// Check Email User Account
	email_search, err := db.Queries.FindEmail(ctx.Background(), result.Email)
	if err != nil {
		return "", errors.New("email not found")
	}

	// Check if Password is same
	if pass.RetypePassword != pass.Password {
		return "", errors.New("password incorect")
	}

	// UpdatePassword
	psw := guard.HashBycrypt(pass.Password) //Hashing Password
	resetPassword := repo.ResetPasswordParams{
		Email:    email_search.Email,
		Password: string(psw),
	}

	if err := db.Queries.ResetPassword(ctx.Background(), resetPassword); err != nil {
		return "", db.Fatal(err)
	}

	// Delete OTP code by email
	if err := rds.DelData(key); err != nil {
		return "", db.Fatal(err)
	}

	return "reset password successfully", nil
}
