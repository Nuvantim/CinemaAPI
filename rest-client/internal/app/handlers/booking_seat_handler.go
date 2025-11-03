// booking handler
package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
	model "booking/pkgs/monorepo"
)

func ListBookingSeat(c *fiber.Ctx) error {
	booking_seat := struct {
		BookingID int64 `json:"booking_id" validate:"required"`
	}{}
	if err := c.BodyParser(&booking_seat); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	if err := validate.BodyStructs(booking_seat);err != nil{
		return c.Status(422).JSON("validate data",err.Error())
	}

	data, err := service.ListBookingSeat(booking_seat)
	if err != nil {
		return c.Status(500).JSON(response.Error("list booking seat", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("list booking seat", data))
}

func CreateBookingSeat(c *fiber.Ctx) error {
	var booking_seat model.CreateBookingSeatParams

	if err := c.BodyParser(&booking_seat); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	data, err := service.CreateBookingSeat(booking_seat)
	if err != nil {
		return c.Status(500).JSON(response.Error("create booking seat", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("create booking seat", data))
}

func DeleteBookingSeat(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("get id", err.Error()))
	}

	if err := service.DeleteBookingSeat(int64(id)); err != nil {
		return c.Status(500).JSON(response.Error("delete booking seat", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("delete booking seat", struct{}{}))
}
