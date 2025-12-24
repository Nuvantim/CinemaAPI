package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func Home(c *fiber.Ctx) error {
	url := os.Getenv("URL")
	link := fmt.Sprintf("%s/docs", url)
	return c.Render("home", fiber.Map{
		"URL": link,
	})
}
