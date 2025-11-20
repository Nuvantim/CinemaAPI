package handler

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
	"os"
)

func Home(c *fiber.Ctx) error {
	url := os.Getenv("URL")
	link := fmt.Sprintf("%s/docs",url)
	return c.Render("home", fiber.Map{
		"URL": link,
	})
}
