package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	model "cinema/pkg/monorepo"
)

func ListScreen(c *fiber.Ctx) error {
	// start service
	data, err := service.ListScreen()
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list screen", err.Error()))
	}
	// return json data
	return c.Status(200).JSON(response.Pass("list screen", data))
}

func GetScreen(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	data, err := service.GetScreen(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get screen", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("success get screen", data))
}
func CreateScreen(c *fiber.Ctx) error {
	// declared model
	var screen model.CreateScreenParams

	// parser body data to json
	if err := c.BodyParser(&screen); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	// validate json
	if err := validate.BodyStructs(screen); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.CreateScreen(screen)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create screen", err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("success create screen", data))
}
func UpdateScreen(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// declared model
	var screen model.UpdateScreenParams

	// parser body data to json
	if err := c.BodyParser(&screen); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}
	// add id to data
	screen.ID = int64(id)

	// validate json
	if err := validate.BodyStructs(screen); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.UpdateScreen(screen)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed update screen", err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("success update screen", data))
}
func DeleteScreen(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	if err := service.DeleteScreen(int64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed delete screen", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("delete screen", struct{}{}))
}
