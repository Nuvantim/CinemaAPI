package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	rds "api/redis"
	model "cinema/pkg/monorepo"
)

func ListCinema(c *fiber.Ctx) error {
	// check data on redis
	redis_data, err := rds.GetData[[]model.Cinema]("list:cinema")
	if err == nil && redis_data != nil {
		c.Status(200).JSON(response.Pass("list cinema", redis_data))
	}
	// start service
	data, err := service.ListCinema()
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list cinema", err.Error()))
	}
	// return json data
	return c.Status(200).JSON(response.Pass("list cinema", data))
}

func ListCinemaSchedule(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}
	// start service
	data, err := service.ListCinemaSchedule(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list cinema schedule", err.Error()))
	}
	// return json data
	return c.Status(200).JSON(response.Pass("list cinema schedule", data))
}

func GetCinema(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	data, err := service.GetCinema(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get cinema", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("success get cinema", data))
}
func CreateCinema(c *fiber.Ctx) error {
	// declared model
	var cinema model.CreateCinemaParams

	// parser body data to json
	if err := c.BodyParser(&cinema); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	// validate json
	if err := validate.BodyStructs(cinema); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.CreateCinema(cinema)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create cinema", err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("success create cinema", data))
}
func UpdateCinema(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// declared model
	var cinema model.UpdateCinemaParams

	// parser body data to json
	if err := c.BodyParser(&cinema); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}
	// add id to data
	cinema.ID = int64(id)

	// validate json
	if err := validate.BodyStructs(cinema); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.UpdateCinema(cinema)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed update cinema", err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("success update cinema", data))
}
func DeleteCinema(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	if err := service.DeleteCinema(int64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed delete cinema", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("delete cinema", struct{}{}))
}
