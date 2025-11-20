// booking handler
package handler

import (
	"github.com/gofiber/fiber/v2"

	req "api/internal/app/request"
	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"

	model "booking/pkg/monorepo"
)

func ListBooking(c *fiber.Ctx) error {
	// get user id
	user_id, ok := c.Locals("user_id").(int64)
	if !ok || user_id == 0 {
		return c.Status(401).JSON(response.Error("failed get user_id", "unauthorized"))
	}

	var body = req.BookingPayment{
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
	user_id, ok := c.Locals("user_id").(int64)
	if !ok || user_id == 0 {
		return c.Status(401).JSON(response.Error("failed get user_id", "unauthorized"))
	}

	// parser json body data
	var booking model.CreateBookingParams
	if err := c.BodyParser(&booking); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}
	booking.UserID = user_id

	// validate data
	if err := validate.BodyStructs(booking); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// check showtime
	_, err := service.GetShowTime(booking.ShowtimeID)
	if err != nil {
		return c.Status(404).JSON(response.Error("failed check showtime", err.Error()))
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
	return c.Status(200).JSON(response.Pass("booking deleted", struct{}{}))
}
