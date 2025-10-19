package routes

import (
	"cinema/config"
	"cinema/internal/app/handlers"
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
	r.Get("/film/genre/", handler.SearchFilmGenre)
	r.Post("/film/create", handler.CreateFilm)
	r.Put("/film/update/", handler.UpdateFilm)
	r.Delete("/film/delete/", handler.DeleteFilm)

	// cinema
	r.Get("/cinemas", handler.ListCinema)
	r.Get("/cinema/", handler.GetCinema)
	r.Post("/cinema/create", handler.CreateCinema)
	r.Put("/cinema/update/", handler.UpdateCinema)
	r.Delete("/cinema/delete/", handler.DeleteCinema)

	// screen type
	r.Get("/screen/types", handler.ListScreenType)
	r.Get("/screen/type/", handler.GetScreenType)
	r.Post("/screen/type/create", handler.CreateScreenType)
	r.Put("/screen/type/update/", handler.UpdateScreenType)
	r.Delete("/screen/type/delete/", handler.DeleteScreenType)

	// screen
	r.Get("/screens", handler.ListScreen)
	r.Get("/screen/", handler.GetScreen)
	r.Post("/screen/create", handler.CreateScreen)
	r.Put("/screen/update/", handler.UpdateScreen)
	r.Delete("/screen/delete/", handler.DeleteScreen)

}
