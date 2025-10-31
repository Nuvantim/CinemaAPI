package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
	model "cinema/pkgs/monorepo"
)

func ListGenre(c *fiber.Ctx) error {
	data, err := service.ListGenre()
	if err != nil {
		return c.Status(500).JSON(response.Error("list genre", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("list genre",data))
}

func GetGenre(c *fiber.Ctx) error {
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("get id", err.Error())
	}

	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(500).JSON(response.Error("validation", err.Error()))
	}

	data, err := service.GetGenre(id)
	if err != nil {
		return c.Status(500).JSON(response.Error("get genre", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("get genre",data))
}

func CreateGenre(c *fiber.Ctx) error {
	var genre model.Genre

	if err := c.BodyParser(&genre); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	if err := validate.BodyStructs(genre); err != nil {
		return c.Status(400).JSON(response.Error("validate data", err.Error()))
	}

	data, err := service.CreateGenre(genre)
	if err != nil {
		return c.Status(500).JSON(response.Error("create genre", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("create genre", data))
}

func UpdateGenre(c *fiber.Ctx) error {
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("get id", err.Error())
	}

	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(500).JSON(response.Error("validation", err.Error()))
	}

	var genre model.Genre
	if err := c.BodyParser(&genre); err != nil {
		return c.Status(400).JSON("parser json", err.Error())
	}

	genre.ID = id

	if err := validate.BodyStructs(genre); err != nil {
		return c.Status(400).JSON("validate data", err.Error())
	}

	data, err := service.UpdateGenre(genre)
	if err != nil {
		return c.Status(500).JSON("update genre", err.Error())
	}

	return c.Status(200).JSON(response.Pass("update genre", data))
}

func DeleteGenre(c *fiber.Ctx) error {
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("get id", err.Error()))
	}
	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(500).JSON(response.Error("validation", err.Error()))
	}

	if err := service.DeleteGenre(id); err != nil {
		return c.Status(500).JSON(response.Error("delete genre", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("delete genre", struct{}{}))
}
