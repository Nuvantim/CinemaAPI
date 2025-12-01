package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	var id = c.Locals("user_id").(int64)
	if id == 0 {
		return c.Status(401).JSON(response.Error("failed get user_id", "unauthorized"))
	}
	// Get Account by id
	user, err := service.GetProfile(id)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get profile", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("success get profile", user))
}

func UpdateAccount(c *fiber.Ctx) error {
	var id = c.Locals("user_id").(int64)
	if id == 0 {
		return c.Status(401).JSON(response.Error("failed get user_id", "unauthorized"))
	}
	var user request.UpdateAccount
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(user); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	userUpdate, err := service.UpdateAccount(user, id)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed update account", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("success update account", userUpdate))
}

func DeleteAccount(c *fiber.Ctx) error {
	var id = c.Locals("user_id").(int64)
	if id == 0 {
		return c.Status(401).JSON(response.Error("failed get user_id", "unauthorized"))
	}
	msg, err := service.DeleteAccount(id)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed delete account", err.Error()))
	}
	return c.Status(200).JSON(response.Pass(msg, struct{}{}))
}
