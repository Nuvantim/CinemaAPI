package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/yokeTH/gofiber-scalar/scalar/v2"
)

func AllowDoc(c *fiber.Ctx) error {
	// Get Server Config
	serverConfig, err := GetServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(c.Path(), "/api/v1/docs") {
		return c.Next()
	}

	return helmet.New(helmet.Config{
		ContentSecurityPolicy: fmt.Sprintf("dafault-src 'self'; frame-ancestors 'none'; http://%s, https://%s", serverConfig.Url, serverConfig.Url),
		HSTSMaxAge:            31536000,
		XFrameOptions:         "SAMEORIGIN",
		HSTSPreloadEnabled:    true,
		HSTSExcludeSubdomains: false,
	})(c)
}

func APIDocs() scalar.Config {
	data, err := os.ReadFile("docs/openapi.json")
	if err != nil {
		log.Fatal("failed read openapi.json :", err)
	}

	var endpoint string = fmt.Sprintf(os.Getenv("URL") + "/api/v1")
	fileContent := strings.ReplaceAll(string(data), "{{BASE_URL}}", endpoint)
	return scalar.Config{
		Title:             "Cinema API Docs",
		BasePath:          "/api/v1",
		Path:              "docs",
		FileContentString: fileContent,
		Theme:             scalar.ThemeBluePlanet,
	}
}
