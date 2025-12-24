package main

import (
	"fmt"
	"log"
	"net/http"

	"cinema/config"
	"cinema/database"
	"cinema/internal/server/http"
	rds "cinema/redis"
)

func main() {
	// Check environment
	if err := config.CheckEnv(); err != nil {
		log.Fatal(err)
	}
	log.Println("Environtment detected....")
	//Start Server
	app, port := server.StartServer()

	// show banner
	config.Banner()

	// Jalankan server di goroutine agar bisa graceful shutdown
	go func() {
		fmt.Println("Server berjalan di http://localhost:" + port + " (HTTP/2 H2C enabled)")
		config.PrintLine()
		if err := app.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Run GraceFullyShutdown
	config.GraceFullShutdown(app)
	// Close Redis
	rds.RedisClose()

	// Close Database
	defer database.CloseDB()

}
