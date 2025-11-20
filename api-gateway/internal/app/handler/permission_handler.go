package handler

import (
	model "api/internal/app/repository"
	"api/internal/app/request"
	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	"github.com/gofiber/fiber/v2"
)

func GetPermission(c *fiber.Ctx) error {
	// Get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	permission, err := service.GetPermission(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get permission", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("success get permission", permission))
}

func ListPermission(c *fiber.Ctx) error {
	permission, err := service.ListPermission()
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list permission", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("list permission", permission))
}

func CreatePermission(c *fiber.Ctx) error {
	var data request.Permission
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(data); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	permission, err := service.CreatePermission(data)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create permission", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("success create permission", permission))
}

func UpdatePermission(c *fiber.Ctx) error {
	var data model.UpdatePermissionParams
	// Get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}
	data.ID = int64(id)

	// validation data
	if err := validate.BodyStructs(data); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	permission, err := service.UpdatePermission(data)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed update permission", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("success update permission", permission))
}
func DeletePermission(c *fiber.Ctx) error {
	// Get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	message, err := service.DeletePermission(int64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed delete permission", err.Error()))
	}

	return c.Status(200).JSON(response.Pass(message, struct{}{}))
}
