// booking handler
package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
	rds "api/redis"

	model "booking/pkgs/monorepo"
)

func ListBooking(c *fiber.Ctx) error {
	// get user id
	user_id, ok := c.Locals("user_id").(int64)
	if !ok || user_id == 0 {
		return c.Status(401).JSON(response.Error("failed get user_id", "unauthorized"))
	}

	// check data on redis
	value := fmt.Sprintf("list:booking:%d", user_id)
	redis_data, err := rds.GetData[[]service.ListBookingRow](value)
	if err == nil && redis_data != nil {
		c.Status(200).JSON(response.Pass("list film", redis_data))
	}

	// create struct data
	body := struct {
		UserID int64 `json:"user_id"`
	}{
		UserID: user_id,
	}

	// start service
	data, err := service.ListBooking(body)
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list booking", err.Error()))
	}

	// return response data
	return c.Status(200).JSON(response.Pass("list booking", data))
}

func CreateBooking(c *fiber.Ctx) error {
	// get user id
	userID, ok := c.Locals("user_id").(int64)
	if !ok || userID == 0 {
		return c.Status(401).JSON(response.Error("failed get user_id", "unauthorized"))
	}

	// parser json body data
	var booking model.CreateBookingParams
	if err := c.BodyParser(&booking); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}
	booking.UserID = userID

	// validate data
	if err := validate.BodyStructs(booking); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.CreateBooking(booking)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create booking", err.Error()))
	}

	// return response data
	return c.Status(200).JSON(response.Pass("success create booking", data))
}

func DeleteBooking(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	if err := service.DeleteBooking(int64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed delete booking", err.Error()))
	}

	// return response
	return c.Status(200).JSON(response.Pass("delete booking", struct{}{}))
}
