// payment handler
package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkgs/utils/responses"
	"api/pkgs/utils/validates"

	model "booking/pkgs/monorepo"
)

func ListPayment(c *fiber.Ctx) error {
	// get user id
	user_id, ok := c.Locals("user_id").(int64)
	if !ok || user_id == 0 {
		return c.Status(401).JSON(response.Error("get user id", "unauthorized"))
	}
	//  add user id to struct
	payment := struct {
		UserID int64 `json:"user_id"`
	}{
		UserID: user_id,
	}

	// start service
	data, err := service.ListPayment(payment)
	if err != nil {
		return c.Status(500).JSON(response.Error("list payment", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("list payment", data))
}

func CreatePayment(c *fiber.Ctx) error {
	// declared model
	var payment model.CreatePaymentParams

	// parser body data to json
	if err := c.BodyParser(&payment); err != nil {
		return c.Status(400).JSON(response.Error("parser json", err.Error()))
	}
	payment.PaymentStatus.String = "Success"

	// validate data
	if err := validate.BodyStructs(payment); err != nil {
		return c.Status(422).JSON(response.Error("validate data", err.Error()))
	}


	// start service
	data, err := service.CreatePayment(payment)
	if err != nil {
		return c.Status(500).JSON(response.Error("create payment", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("create payment", data))
}
