package routes

import (
	"cinema/config"
	"cinema/internal/app/handler"
	"net/http"
)

func Setup(mux *http.ServeMux) {
	r := &config.Router{Mux: mux}
	r.Get("/", handler.GetTest)

	// genre
	r.Get("/genres", handler.ListGenre)
	r.Get("/genre/", handler.GetGenre)
	r.Post("/genre/create", handler.CreateGenre)
	r.Put("/genre/update", handler.UpdateGenre)
	r.Delete("/genre/delete/", handler.DeleteGenre)

	// film
	r.Get("/films", handler.ListFilm)
	r.Get("/film/", handler.GetFilm)
	r.Post("/film/search", handler.SearchFilm)
	r.Get("/film/genre/", handler.SearchFilmGenre)
	r.Post("/film/create", handler.CreateFilm)
	r.Put("/film/update", handler.UpdateFilm)
	r.Delete("/film/delete/", handler.DeleteFilm)

	// cinema
	r.Get("/cinemas", handler.ListCinema)
	r.Get("/cinema/schedules/", handler.ListCinemaSchedule)
	r.Get("/cinema/", handler.GetCinema)
	r.Post("/cinema/create", handler.CreateCinema)
	r.Put("/cinema/update", handler.UpdateCinema)
	r.Delete("/cinema/delete/", handler.DeleteCinema)

	// screen type
	r.Get("/screen/types", handler.ListScreenType)
	r.Get("/screen/type/", handler.GetScreenType)
	r.Post("/screen/type/create", handler.CreateScreenType)
	r.Put("/screen/type/update", handler.UpdateScreenType)
	r.Delete("/screen/type/delete/", handler.DeleteScreenType)

	// screen
	r.Get("/screens", handler.ListScreen)
	r.Get("/screen/", handler.GetScreen)
	r.Post("/screen/create", handler.CreateScreen)
	r.Put("/screen/update", handler.UpdateScreen)
	r.Delete("/screen/delete/", handler.DeleteScreen)

	// seat
	r.Get("/seats", handler.ListSeat)
	r.Get("/seat/", handler.GetSeat)
	r.Post("/seat/create", handler.CreateSeat)
	r.Post("/seat/price", handler.SeatPrice)
	r.Put("/seat/update", handler.UpdateSeat)
	r.Delete("/seat/delete/", handler.DeleteSeat)

	// showtime
	r.Get("/showtimes", handler.ListShowTime)
	r.Get("/showtime/", handler.GetShowTime)
	r.Post("/showtime/create", handler.CreateShowTime)
	r.Put("/showtime/update", handler.UpdateShowTime)
	r.Delete("/showtime/delete/", handler.DeleteShowTime)

}
