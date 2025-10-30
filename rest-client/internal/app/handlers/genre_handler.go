package handler

import (
	"github.com/gofiber/fiber/v2"

	model "cinema/pkgs/monorepo"
	resp "api/pkgs/utils/responses"
	val "api/pkgs/utils/validates"
	"api/internal/app/service"
)

func ListGenre(c *fiber.Ctx) error {
	data, err := service.ListGenre()
	if err != nil {
		return c.Status(500).JSON(resp.Error("list genre", err.Error()))
	}

	return c.Status(200).JSON(data)
}

func GetGenre(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("parser id", err.Error())
	}

	data, err := service.GetGenre(int32(id))
	if err != nil {
		return c.Status(500).JSON(resp.Error("get genre", err.Error()))
	}

	return c.Status(200).JSON(data)
}

func CreateGenre(c *fiber.Ctx) error {
	var genre model.Genre

	if err := c.BodyParser(&genre); err != nil {
		return c.Status(400).JSON(resp.Error("parser json", err.Error()))
	}

	if err := val.BodyStructs(genre); err != nil {
		return c.Status(400).JSON(resp.Error("validate data", err.Error()))
	}

	data, err := service.CreateGenre(genre)
	if err != nil {
		return c.Status(500).JSON(resp.Error("create genre", err.Error()))
	}

	return c.Status(200).JSON(resp.Pass("create genre", data))
}

func UpdateGenre(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("parser id", err.Error())
	}

	var genre model.Genre
	if err := c.BodyParser(&genre); err != nil {
		return c.Status(400).JSON("parser json", err.Error())
	}

	genre.ID = int32(id)

	if err := val.BodyStructs(genre); err != nil {
		return c.Status(400).JSON("validate data", err.Error())
	}

	data, err := service.UpdateGenre(genre)
	if err != nil {
		return c.Status(500).JSON("update genre", err.Error())
	}

	return c.Status(200).JSON(resp.Pass("update data", data))
}

func DeleteGenre(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(resp.Error("parser id", err.Error()))
	}

	if err := service.DeleteGenre(int32(id)); err != nil {
		return c.Status(500).JSON(resp.Error("delete genre", err.Error()))
	}

	return c.Status(200).JSON(resp.Pass("delete genre", struct{}{}))
}
