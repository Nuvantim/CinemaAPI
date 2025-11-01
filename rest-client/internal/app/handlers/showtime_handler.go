package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
	model "cinema/pkgs/monorepo"
)

func ListShowTime(c *fiber.Ctx) error {
	// start service
	data, err := service.ListShowTime()
	if err != nil {
		return c.Status(500).JSON(response.Error("list showtime", err.Error()))
	}
	// return json data
	return c.Status(200).JSON(response.Pass("list showtime", data))
}

func GetShowTime(c *fiber.Ctx) error {
	// get params id
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("get id", err.Error()))
	}

	// id validation
	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(400).JSON(response.Error("validation", err.Error()))
	}

	// start service
	data, err := service.GetShowTime(id)
	if err != nil {
		return c.Status(500).JSON(response.Error("get showtime", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("get showtime", data))
}
func CreateShowTime(c *fiber.Ctx) error {
	// declared model
	var showtime model.CreateShowTimeParams

	// parser body data to json
	if err := c.BodyParser(&showtime); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	// validate json
	if err := validate.BodyStructs(showtime); err != nil {
		return c.Status(422).JSON(response.Error("validate data", err.Error()))
	}

	// start service
	data, err := service.CreateShowTime(showtime)
	if err != nil {
		return c.Status(500).JSON(response.Error("create showtime", err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("create showtime", data))
}
func UpdateShowTime(c *fiber.Ctx) error {
	// get params id
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("get id", err.Error()))
	}

	// id validation
	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(400).JSON(response.Error("valdation", err.Error()))
	}

	// declared model
	var showtime model.UpdateShowTimeParams

	// parser body data to json
	if err := c.BodyParser(&showtime); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}
	// add id to data
	showtime.ID = id

	// validate json
	if err := validate.BodyStructs(showtime); err != nil {
		return c.Status(422).JSON(response.Error("validate data", err.Error()))
	}

	// start service
	data, err := service.UpdateShowTime(showtime)
	if err != nil {
		return c.Status(500).JSON(response.Error("update showtime", err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("update showtime", data))
}
func DeleteShowTime(c *fiber.Ctx) error {
	// get params id
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("get id", err.Error()))
	}

	// id validation
	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(400).JSON(response.Error("validation", err.Error()))
	}

	// start service
	if err := service.DeleteShowTime(id); err != nil {
		return c.Status(500).JSON(response.Error("delete showtime", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("delete showtime", struct{}{}))
}
