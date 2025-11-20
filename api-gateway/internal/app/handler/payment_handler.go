// payment handler
package handler

import (
	"github.com/gofiber/fiber/v2"

	req "api/internal/app/request"
	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"

	model "booking/pkg/monorepo"
)

func ListPayment(c *fiber.Ctx) error {
	// get user id
	user_id, ok := c.Locals("user_id").(int64)
	if !ok || user_id == 0 {
		return c.Status(401).JSON(response.Error("failed get user id", "unauthorized"))
	}

	//  add user id to struct
	var payment = req.BookingPayment{
		UserID: user_id,
	}

	// start service
	data, err := service.ListPayment(payment)
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list payment", err.Error()))
	}

	// return json data
	return c.Status(200).JSON(response.Pass("list payment", data))
}

func CreatePayment(c *fiber.Ctx) error {
	// declared model
	var payment model.CreatePaymentParams

	// parser body data to json
	if err := c.BodyParser(&payment); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	// input payment status
	payment.PaymentStatus = "Success"

	// validate data
	if err := validate.BodyStructs(payment); err != nil {
		return c.Status(422).JSON(response.Error("invalid or incomplete data", err.Error()))
	}

	// start service
	data, err := service.CreatePayment(payment)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create payment", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("success create payment", data))
}
