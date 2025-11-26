package server

import (
	"log"
	"net/http"
	"time"

	"cinema/config"
	"cinema/database"
	"cinema/internal/routes"
	rds "cinema/redis"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func StartServer() (*http.Server, string) {
	// Initialization Server
	mux := http.NewServeMux()

	// Set Routes
	routes.Setup(mux)

	// Set Log Request
	handler := config.LoggingRequest(mux)

	// Database Connection
	database.InitDB()

	// Redis Connection
	rds.InitRedis()

	// Get Server Configuration
	conf, err := config.GetServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	// http2 configuration
	h2s := &http2.Server{}
	server := &http.Server{
		Addr:              ":" + conf.Port,
		Handler:           h2c.NewHandler(handler, h2s),
		ReadHeaderTimeout: 5 * time.Second,
	}

	return server, conf.Port
}
