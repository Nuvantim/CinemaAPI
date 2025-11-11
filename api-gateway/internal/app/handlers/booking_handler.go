// booking handler
package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"

	model "booking/pkgs/monorepo"
)

func ListBooking(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(int64)
	if !ok || userID == 0 {
		return c.Status(401).JSON(response.Error("failed get user_id", "unauthorized"))
	}

	body := struct {
		UserID int64 `json:"user_id"`
	}{
		UserID: userID,
	}

	data, err := service.ListBooking(body)
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list booking", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("list booking", data))
}

func CreateBooking(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(int64)
	if !ok || userID == 0 {
		return c.Status(401).JSON(response.Error("failed get user_id", "unauthorized"))
	}

	var booking model.CreateBookingParams
	if err := c.BodyParser(&booking); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}
	booking.UserID = userID

	if err := validate.BodyStructs(booking); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	data, err := service.CreateBooking(booking)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create booking", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("success create booking", data))
}

func DeleteBooking(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	if err := service.DeleteBooking(int64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed delete booking", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("delete booking", struct{}{}))
}
