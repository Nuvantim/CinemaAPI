package handler

import (
	"api/config"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{})
}

func Swagger(c *fiber.Ctx) error {
	var url string
	serv, _ := config.GetServerConfig()
	url = serv.Url
	if url == "example.com" {
		url = "127.0.0.1" + ":" + serv.Port
	}
	return c.Render("api", fiber.Map{
		"Title":      "CinemaAPI Documentation",
		"APIBaseURL": url,
	})
}
