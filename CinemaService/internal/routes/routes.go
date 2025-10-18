package routes

import (
	"cinema/config"
	"cinema/internal/app/handler"
	"net/http"
)

func Setup(mux *http.ServeMux) {
	r := &config.Router{mux}
	r.Get("/", handler.GetTest)

	// genre
	r.Get("/genres", handler.ListGenre)
	r.Get("/genre/", handler.GetGenre)
	r.Post("/genre/create", handler.CreateGenre)
	r.Put("/genre/update/", handler.UpdateGenre)
	r.Delete("/genre/delete/", handler.DeleteGenre)

	// film
	r.Get("/films", handler.ListGenre)
	r.Get("/film/", handler.GetGenre)
	r.Get("/film/search", handler.SearchFilm)
	r.Get("/film/genre/", handler.SearchGenreFilm)
	r.Post("/film/create", handler.CreateFilm)
	r.Put("/film/update/", handler.UpdateFilm)
	r.Delete("/film/delete/", handler.DeleteFilm)
}
