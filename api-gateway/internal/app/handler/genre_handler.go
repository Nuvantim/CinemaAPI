package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	model "cinema/pkg/monorepo"
)

func ListGenre(c *fiber.Ctx) error {
	// start service
	data, err := service.ListGenre()
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list genre", err.Error()))
	}
	// return json data
	return c.Status(200).JSON(response.Pass("list genre", data))
}

func GetGenre(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("get id", err.Error())
	}

	// start service
	data, err := service.GetGenre(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get genre", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("success get genre", data))
}

func CreateGenre(c *fiber.Ctx) error {
	// declared model
	var genre = struct {
		Name string `json:"name" validate:"required"`
	}{}

	// parser body data to json
	if err := c.BodyParser(&genre); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	// validate json
	if err := validate.BodyStructs(genre); err != nil {
		return c.Status(400).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.CreateGenre(genre)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create genre", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("success create genre", data))
}

func UpdateGenre(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("get id", err.Error())
	}

	// declared mode
	var genre model.Genre

	// parser body data to json
	if err := c.BodyParser(&genre); err != nil {
		return c.Status(400).JSON("parser json", err.Error())
	}

	// add id to data
	genre.ID = int64(id)

	// validate data
	if err := validate.BodyStructs(genre); err != nil {
		return c.Status(400).JSON("validate data", err.Error())
	}

	// start service
	data, err := service.UpdateGenre(genre)
	if err != nil {
		return c.Status(500).JSON("update genre", err.Error())
	}

	// return json data
	return c.Status(200).JSON(response.Pass("success update genre", data))
}

func DeleteGenre(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	if err := service.DeleteGenre(int64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed delete genre", err.Error()))
	}

	// return response
	return c.Status(200).JSON(response.Pass("delete genre", struct{}{}))
}
