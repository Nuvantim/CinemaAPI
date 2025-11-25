package config

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/yokeTH/gofiber-scalar/scalar/v2"
)

func AllowDoc(c *fiber.Ctx) error {
	if strings.HasPrefix(c.Path(), "/api/v1/docs") {
		return c.Next()
	}
	return helmet.New(helmet.Config{
		ContentSecurityPolicy: "dafault-src 'self'; frame-ancestors 'self'",
		HSTSMaxAge:            31536000,
		HSTSPreloadEnabled:    true,
		HSTSExcludeSubdomains: false,
	})(c)
}

func APIDocs() scalar.Config {
	data, err := os.ReadFile("docs/openapi.json")
	if err != nil {
		log.Fatal("failed read openapi.json :", err)
	}
	fileContent := strings.ReplaceAll(string(data), "{{BASE_URL}}", os.Getenv("URL"))
	return scalar.Config{
		Title:             "Cinema API Docs",
		BasePath:          "/api/v1",
		Path:              "docs",
		FileContentString: fileContent,
		Theme:             scalar.ThemeBluePlanet,
	}
}
