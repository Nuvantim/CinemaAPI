package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	"github.com/gofiber/fiber/v2"
)

func GetClient(c *fiber.Ctx) error {
	// Get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	client, err := service.GetClient(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get client data", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("success get client data", client))
}

func ListClient(c *fiber.Ctx) error {
	client, err := service.ListClient()
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list client", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("list client", client))
}

func UpdateClient(c *fiber.Ctx) error {
	// Get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	var data request.UpdateClient

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(data); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	client, err := service.UpdateClient(int64(id), data)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed update client", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("success update client", client))
}

func DeleteClient(c *fiber.Ctx) error {
	// Get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	message, err := service.DeleteClient(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed delete client", err.Error()))
	}

	return c.Status(200).JSON(response.Pass(message, struct{}{}))
}
