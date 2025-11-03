// payment handler
package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkgs/utils/responses"

	model "cinema/pkgs/monorepo"
)

func ListPayment(c *fiber.Ctx) error {
	user_id, ok := c.Locals("user_id").(int64)
	if !ok || user_id == 0 {
		return c.Status(401).JSON(response.Error("get user id", "unauthorized"))
	}

	payment := struct {
		UserID int64 `json:"user_id"`
	}{
		UserID: user_id,
	}

	data, err := service.ListPayment(payment)
	if err != nil {
		return c.Status(500).JSON(response.Error("list payment", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("list payment", data))
}

func CreatePayment(c *fiber.Ctx) error {
	var payment model.CreatePaymentParams

	if err := c.BodyParser(&payment); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}

	data, err := service.CreatePayment(payment)
	if err != nil {
		return c.Status(500).JSON(response.Error("create payment", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("create payment", data))
}
