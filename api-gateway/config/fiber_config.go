package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func FiberConfig() fiber.Config {
	// set folder views as html rendering
	engine := html.New("./views", ".html")

	// Get Environtmet Server Config
	var envServ, err = GetServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	// return fiber config
	return fiber.Config{
		AppName:       envServ.AppName,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Nuvantim Project",
		Prefork:       false,
		Views:         engine,
	}
}

func SecurityConfig(app *fiber.App) {

	// Rate Limiting
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 120 * time.Second,
	}))

	// Logger
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	//Idempotency
	app.Use(idempotency.New())

	// CSRF Protection
	// app.Use(csrf.New())

	// CORS Configuration
	var url string = os.Getenv("URL")
	var port string = os.Getenv("PORT")

	var origin = fmt.Sprintf("%s,http://localhost:%s, http://127.0.0.1:%s", url, port, port)
	app.Use(cors.New(cors.Config{
		AllowOrigins: origin,
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Authorization, Accept",
		MaxAge:       3600,
	}))
}
