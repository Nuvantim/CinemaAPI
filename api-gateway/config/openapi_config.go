package config

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/yokeTH/gofiber-scalar/scalar/v2"
)

func APIDocs() scalar.Config {
	root, err := os.OpenRoot("docs")
	if err != nil {
		log.Fatalf("error membuka folder docs: %w", err)
	}
	defer root.Close()
	file, err := root.Open("openapi.yaml")
	if err != nil {
		log.Fatalf("error membuka file openapi.yaml: %w", err)
	}
	defer file.Close()
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		log.Fatalf("gagal baca template: %v", err)
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
