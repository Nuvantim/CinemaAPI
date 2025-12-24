package handler

import (
	"api/internal/app/request"
	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	var data request.Role
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(data); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	role, err := service.CreateRole(data)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create role", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("success create role", role))
}
func GetRole(c *fiber.Ctx) error {
	// Get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	role, err := service.GetRole(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get role", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("success get role", role))
}
func ListRole(c *fiber.Ctx) error {
	role, err := service.ListRole()
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list role", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("list role", role))
}
func UpdateRole(c *fiber.Ctx) error {
	var data request.Role
	// Get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// parser json
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(data); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	role, err := service.UpdateRole(data, int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed update role", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("success update role", role))

}
func DeleteRole(c *fiber.Ctx) error {
	// Get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	msg, err := service.DeleteRole(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed delete role", err.Error()))
	}
	return c.Status(200).JSON(response.Pass(msg, struct{}{}))
}
