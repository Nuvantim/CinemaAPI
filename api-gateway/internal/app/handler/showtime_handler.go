package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	rds "api/redis"
	model "cinema/pkg/monorepo"
)

func ListShowTime(c *fiber.Ctx) error {
	// check data on redis
	redis_data, err := rds.GetData[[]model.ListShowTimeRow]("list:showtime")
	if err == nil && redis_data != nil {
		c.Status(200).JSON(response.Pass("list showtime", redis_data))
	}
	// start service
	data, err := service.ListShowTime()
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list showtime", err.Error()))
	}
	// return json data
	return c.Status(200).JSON(response.Pass("list showtime", data))
}

func GetShowTime(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	data, err := service.GetShowTime(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get showtime", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("success get showtime", data))
}
func CreateShowTime(c *fiber.Ctx) error {
	// declared model
	var showtime model.CreateShowTimeParams

	// parser body data to json
	if err := c.BodyParser(&showtime); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	// validate json
	if err := validate.BodyStructs(showtime); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.CreateShowTime(showtime)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create showtime", err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("success create showtime", data))
}
func UpdateShowTime(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// declared model
	var showtime model.UpdateShowTimeParams

	// parser body data to json
	if err := c.BodyParser(&showtime); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}
	// add id to data
	showtime.ID = int64(id)

	// validate json
	if err := validate.BodyStructs(showtime); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.UpdateShowTime(showtime)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed update showtime", err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("success update showtime", data))
}
func DeleteShowTime(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	if err := service.DeleteShowTime(int64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed delete showtime", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("delete showtime", struct{}{}))
}
