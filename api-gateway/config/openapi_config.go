package config

import (
	"log"
	"os"
	"strings"

	"github.com/yokeTH/gofiber-scalar/scalar/v2"
)

func APIDocs() scalar.Config {
	data, err := os.ReadFile("docs/openapi.json")
	if err != nil {
		log.Fatalf("failed read file openapi.json: %v", err)
	}

	fileContent := strings.ReplaceAll(string(data), "{{BASE_URL}}", os.Getenv("URL"))

	return scalar.Config{
		Title:             "Cinema API Docs",
		BasePath:          "/api/v1",
		Path:              "docs",
		CacheAge:          60,
		FallbackCacheAge:  86400,
		FileContentString: fileContent,
		Theme:             scalar.ThemeBluePlanet,
	}
}
