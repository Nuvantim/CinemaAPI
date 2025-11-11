package config

import (
	"log"
	"os"
	"strings"

	"github.com/yokeTH/gofiber-scalar/scalar/v2"
)

func APIDocs() scalar.Config {
	data, err := os.ReadFile("docs/openapi.yaml")
	if err != nil {
		log.Fatalf("failed read file openapi.yaml: %v", err)
	}

	fileContent := strings.ReplaceAll(string(data), "{{BASE_URL}}", os.Getenv("URL"))

	return scalar.Config{
		Title:             "Cinema API Docs",
		BasePath:          "/",
		Path:              "/api/v1/docs",
		FileContentString: fileContent,
		Theme:             scalar.ThemeBluePlanet,
	}
}
