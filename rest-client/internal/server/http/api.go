package http

import (
	"fmt"

	"api/config"
	"api/database"
	"api/internal/routes"
	// rds "api/redis"
	"api/internal/server/client"

	"github.com/gofiber/fiber/v2"
)

// ServerGo initializes and returns a Fiber app instance
func ServerGo() *fiber.App {
	// Start Fiber APP
	app := fiber.New(config.FiberConfig())

	// Security Configuration
	config.SecurityConfig(app)

	// Set up all routes
	routes.Setup(app)

	// Start Database Connection
	fmt.Print("[ 🗄️ ]")
	database.InitDB()

	// Start redis Connection
	// fmt.Print("[ 🎲 ]")
	// rds.InitRedis()

	// Star Connect Cinema Service
	fmt.Print("[ 🎬 ]")
	client.CinemaService()

	return app
}
