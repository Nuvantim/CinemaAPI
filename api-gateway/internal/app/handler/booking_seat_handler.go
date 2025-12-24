// booking handler
package handler

import (
	"github.com/gofiber/fiber/v2"

	"api/internal/app/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	model "booking/pkg/monorepo"
)

func ListBookingSeat(c *fiber.Ctx) error {
	booking_seat := struct {
		BookingID int64 `json:"booking_id" validate:"required"`
	}{}
	if err := c.BodyParser(&booking_seat); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	if err := validate.BodyStructs(booking_seat); err != nil {
		return c.Status(422).JSON("validate data", err.Error())
	}

	data, err := service.ListBookingSeat(booking_seat)
	if err != nil {
		return c.Status(500).JSON(response.Error("unable to list booking seat", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("list booking seat", data))
}

func CreateBookingSeat(c *fiber.Ctx) error {
	var booking_seat model.CreateBookingSeatParams

	if err := c.BodyParser(&booking_seat); err != nil {
		return c.Status(400).JSON(response.Error("unable to parse request body", err.Error()))
	}

	// check seat
	_, err := service.GetSeat(booking_seat.SeatID)
	if err != nil {
		return c.Status(404).JSON(response.Error("failed check seat", err.Error()))
	}

	// get booking
	booking, err := service.GetBooking(booking_seat.BookingID)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get booking data", err.Error()))
	}

	// get price seat
	price, err := service.SeatPrice(booking.ShowtimeID, booking_seat.SeatID)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed get seat price", err.Error()))
	}

	// input price
	booking_seat.PricePaid = float64(price)

	if err := validate.BodyStructs(booking_seat); err != nil {
		return c.Status(422).JSON(response.Error("validates data", err.Error()))
	}

	data, err := service.CreateBookingSeat(booking_seat)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed create booking seat", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("success create booking seat", data))
}

func DeleteBookingSeat(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := service.DeleteBookingSeat(id); err != nil {
		return c.Status(500).JSON(response.Error("failed delete booking seat", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("delete booking seat", struct{}{}))
}
