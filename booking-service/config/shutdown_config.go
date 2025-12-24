package config

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func GraceFullShutdown(server *http.Server) {
	// Tangkap sinyal Ctrl+C (interrupt)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	fmt.Println("\nMenerima sinyal shutdown...")

	// Graceful shutdown dengan timeout 5 detik
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown gagal: %v", err)
	}

	fmt.Println("Server berhenti dengan aman ✅")
}
