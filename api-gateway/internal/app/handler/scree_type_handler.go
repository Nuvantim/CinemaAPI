package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	model "cinema/pkg/monorepo"
)

func ListScreenType(c *fiber.Ctx) error {
	// start service
	data, err := service.ListScreenType()
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list screen type", err.Error()))
	}
	// return json data
	return c.Status(200).JSON(response.Pass("list screen type", data))
}

func GetScreenType(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	data, err := service.GetScreenType(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get screen type", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("success get screen type", data))
}
func CreateScreenType(c *fiber.Ctx) error {
	// declared model
	var screen_type = struct {
		Name string `json:"name" validate:"required"`
	}{}

	// parser body data to json
	if err := c.BodyParser(&screen_type); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	// validate json
	if err := validate.BodyStructs(screen_type); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.CreateScreenType(screen_type)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create screen type", err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("success create screen type", data))
}
func UpdateScreenType(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// declared model
	var screen_type model.ScreenType

	// parser body data to json
	if err := c.BodyParser(&screen_type); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}
	// add id to data
	screen_type.ID = int64(id)

	// validate json
	if err := validate.BodyStructs(screen_type); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.UpdateScreenType(screen_type)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed update screen type", err.Error()))
	}

	// returning json data
	return c.Status(200).JSON(response.Pass("success update screen type", data))
}
func DeleteScreenType(c *fiber.Ctx) error {
	// get params id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	if err := service.DeleteScreenType(int64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed delete screen type", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("delete screen type", struct{}{}))
}
