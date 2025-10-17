package routes

import (
	"cinema/config"
	"cinema/internal/app/handlers"
	"net/http"
)

func Setup(mux *http.ServeMux) {
	r := &config.Router{mux}
	r.Get("/test", handlers.GetTest)

	// genre
	r.Get("/genres", handlers.ListGenre)
	r.Get("/genre/", handlers.GetGenre)
	r.Post("/genre/create", handlers.CreateGenre)
	r.Put("/genre/update", handlers.UpdateGenre)
	r.Delete("/genre/delete/", handlers.DeleteGenre)
}
