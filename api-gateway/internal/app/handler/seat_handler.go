package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	model "cinema/pkg/monorepo"
)

func ListSeat(c *fiber.Ctx) error {
	// start service
	data, err := service.ListSeat()
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list seat", err.Error()))
	}
	// return json data
	return c.Status(200).JSON(response.Pass("list seat", data))
}

func GetSeat(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	data, err := service.GetSeat(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get seat", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("success get seat", data))
}
func CreateSeat(c *fiber.Ctx) error {
	// declared model
	var seat model.CreateSeatParams

	// parser body data to json
	if err := c.BodyParser(&seat); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	// validate json
	if err := validate.BodyStructs(seat); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.CreateSeat(seat)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create seat", err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("success create seat", data))
}
func UpdateSeat(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// declared model
	var seat model.UpdateSeatParams

	// parser body data to json
	if err := c.BodyParser(&seat); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}
	// add id to data
	seat.ID = int64(id)

	// validate json
	if err := validate.BodyStructs(seat); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.UpdateSeat(seat)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed update seat", err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("success update seat", data))
}
func DeleteSeat(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	if err := service.DeleteSeat(int64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed delete seat", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("delete seat", struct{}{}))
}
