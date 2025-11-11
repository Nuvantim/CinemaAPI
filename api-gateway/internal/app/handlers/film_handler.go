package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"
	model "cinema/pkgs/monorepo"
)

func ListFilm(c *fiber.Ctx) error {
	// start service
	data, err := service.ListFilm()
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list film", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("list film", data))
}

func GetFilm(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("get id", err.Error())
	}

	// start service
	data, err := service.GetFilm(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get film", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("success get film", data))
}

func SearchFilm(c *fiber.Ctx) error {
	// declared model
	var film = struct {
		Title string `json:"title" validate:"required"`
	}{}

	// parser body to json
	if err := c.BodyParser(&film); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(film); err != nil {
		return c.Status(400).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.SearchFilm(film)
	if err != nil {
		return c.Status(500).JSON(response.Error("search film", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("search film", data))
}

func SearchFilmGenre(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("get id", err.Error())
	}

	// start service
	data, err := service.SearchFilmGenre(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("search film genre", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("search film genre", data))
}

func CreateFilm(c *fiber.Ctx) error {
	// declared model
	var film model.CreateFilmParams

	// parser body to json
	if err := c.BodyParser(&film); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(film); err != nil {
		return c.Status(400).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.CreateFilm(film)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create film", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("success create film", data))
}

func UpdateFilm(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("get id", err.Error())
	}

	// declared model
	var film model.UpdateFilmParams

	// parser body to json
	if err := c.BodyParser(&film); err != nil {
		return c.Status(400).JSON("parser json", err.Error())
	}

	// add id to data
	film.ID = int64(id)

	// validate data
	if err := validate.BodyStructs(film); err != nil {
		return c.Status(400).JSON("validate data", err.Error())
	}

	// start service
	data, err := service.UpdateFilm(film)
	if err != nil {
		return c.Status(500).JSON("update film", err.Error())
	}

	// return json data
	return c.Status(200).JSON(response.Pass("success update film", data))
}

func DeleteFilm(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	if err := service.DeleteFilm(int64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed delete film", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("delete film", struct{}{}))
}
