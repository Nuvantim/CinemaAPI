package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
	model "cinema/pkgs/monorepo"
)

func ListFilm(c *fiber.Ctx) error {
	data, err := service.ListFilm()
	if err != nil {
		return c.Status(500).JSON(response.Error("list film", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("list film",data))
}

func GetFilm(c *fiber.Ctx) error {
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("get id", err.Error())
	}

	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(500).JSON(response.Error("validation", err.Error()))
	}

	data, err := service.GetFilm(id)
	if err != nil {
		return c.Status(500).JSON(response.Error("get film", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("get film", data))
}

func SearchFilm(c *fiber.Ctx) error {
	var film model.Film

	if err := c.BodyParser(&film); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	if err := validate.BodyStructs(film); err != nil {
		return c.Status(400).JSON(response.Error("validate data", err.Error()))
	}

	data, err := service.SearchFilm(film)
	if err != nil {
		return c.Status(500).JSON(response.Error("create film", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("search film", data))
}

func SearchFilmGenre(c *fiber.Ctx) error {
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("get id", err.Error())
	}

	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(500).JSON(response.Error("validation", err.Error()))
	}

	data, err := service.SearchFilmGenre(id)
	if err != nil {
		return c.Status(500).JSON(response.Error("search film genre", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("search film genre", data))
}

func CreateFilm(c *fiber.Ctx) error {
	var film model.CreateFilmParams

	if err := c.BodyParser(&film); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	if err := validate.BodyStructs(film); err != nil {
		return c.Status(400).JSON(response.Error("validate data", err.Error()))
	}

	data, err := service.CreateFilm(film)
	if err != nil {
		return c.Status(500).JSON(response.Error("create film", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("create film", data))
}

func UpdateFilm(c *fiber.Ctx) error {
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("get id", err.Error())
	}

	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(500).JSON(response.Error("validation", err.Error()))
	}

	var film model.UpdateFilmParams
	if err := c.BodyParser(&film); err != nil {
		return c.Status(400).JSON("parser json", err.Error())
	}

	film.ID = id

	if err := validate.BodyStructs(film); err != nil {
		return c.Status(400).JSON("validate data", err.Error())
	}

	data, err := service.UpdateFilm(film)
	if err != nil {
		return c.Status(500).JSON("update film", err.Error())
	}

	return c.Status(200).JSON(response.Pass("update film", data))
}

func DeleteFilm(c *fiber.Ctx) error {
	params, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("get id", err.Error()))
	}
	id, err := validate.ValID(params)
	if err != nil {
		return c.Status(500).JSON(response.Error("validation", err.Error()))
	}

	if err := service.DeleteFilm(id); err != nil {
		return c.Status(500).JSON(response.Error("delete film", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("delete film", struct{}{}))
}
