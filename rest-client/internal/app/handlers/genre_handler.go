package handler

import(
	"api/internal/app/service"
)


func ListGenre(c *fiber.Ctx) error {
	data, err := service.ListGenre()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return c.Status(200).JSON(data)
}

func GetGenre(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	data, err := service.GetGenre(int32(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	return c.Status(200).JSON(data)
}