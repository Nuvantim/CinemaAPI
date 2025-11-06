package routes

import (
	"api/internal/app/handlers"
	"api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", handler.Home)

	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/send-otp", handler.SendOTP)
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Post("/reset-password", handler.ResetPassword)

	// Set Middleware
	api.Use(middleware.BearerAuth())

	// user
	account := api.Group("/account")
	account.Get("/profile", handler.GetProfile)
	account.Put("/update", handler.UpdateAccount)
	account.Delete("/delete", handler.DeleteAccount)

	// client
	client := api.Group("/client", middleware.Role("admin"))
	client.Get("/", handler.ListClient)
	client.Get("/:id", handler.GetClient)
	client.Put("/update/:id", handler.UpdateClient)
	client.Delete("/delete/:id", handler.DeleteClient)

	// role
	role := api.Group("/role", middleware.Permission("handle role"))
	role.Get("/", handler.ListRole)
	role.Get("/:id", handler.GetRole)
	role.Post("/store", handler.CreateRole)
	role.Put("/update/:id", handler.UpdateRole)
	role.Delete("/delete/:id", handler.DeleteRole)

	// permission
	permission := api.Group("/permission", middleware.Permission("handle permission"))
	permission.Get("/", handler.ListPermission)
	permission.Get("/:id", handler.GetPermission)
	permission.Post("/store", handler.CreatePermission)
	permission.Put("/update/:id", handler.UpdatePermission)
	permission.Delete("/delete/:id", handler.DeletePermission)

	// genre
	genre := api.Group("/genre")
	genre.Get("/", handler.ListGenre)
	genre.Get("/:id", handler.GetGenre)
	genre.Post("/create", handler.CreateGenre)
	genre.Put("/update/:id", handler.UpdateGenre)
	genre.Delete("/delete/:id", handler.DeleteGenre)

	// film
	film := api.Group("/film")
	film.Get("/", handler.ListFilm)
	film.Get("/:id", handler.GetFilm)
	film.Post("/film/search", handler.SearchFilm)
	film.Get("/film/genre/:id", handler.SearchFilmGenre)
	film.Post("/create", handler.CreateFilm)
	film.Put("/update", handler.UpdateFilm)
	film.Delete("/delete/:id", handler.DeleteFilm)

	// cinema
	cinema := api.Group("/cinema")
	cinema.Get("/", handler.ListCinema)
	cinema.Get("/schedule/:id", handler.ListCinemaSchedule)
	cinema.Get("/:id", handler.GetCinema)
	cinema.Post("/create", handler.CreateCinema)
	cinema.Put("/update", handler.UpdateCinema)
	cinema.Delete("/delete/:id", handler.DeleteCinema)

	// screen type
	screen_type := api.Group("/screen/type")
	screen_type.Get("/", handler.ListScreenType)
	screen_type.Get("/:id", handler.GetScreenType)
	screen_type.Post("/create", handler.CreateScreenType)
	screen_type.Put("/update", handler.CreateScreenType)
	screen_type.Delete("/delete/:id", handler.DeleteScreenType)

	// screen
	screen := api.Group("/screen")
	screen.Get("/", handler.ListScreen)
	screen.Get("/:id", handler.GetScreen)
	screen.Post("/create", handler.CreateScreen)
	screen.Put("/update", handler.UpdateScreen)
	screen.Delete("/delete/:id", handler.DeleteScreen)

	// seat
	seat := api.Group("/seat")
	seat.Get("/", handler.ListSeat)
	seat.Get("/:id", handler.GetSeat)
	seat.Post("/create", handler.CreateSeat)
	seat.Put("/update", handler.UpdateSeat)
	seat.Delete("/delete/:id", handler.DeleteSeat)

	// showtime
	showtime := api.Group("/showtime")
	showtime.Get("/", handler.ListShowTime)
	showtime.Get("/:id", handler.GetShowTime)
	showtime.Post("/create", handler.CreateShowTime)
	showtime.Put("/update", handler.UpdateShowTime)
	showtime.Delete("/delete/:id", handler.DeleteShowTime)

	// booking
	booking := api.Group("/booking")
	booking.Get("/", handler.ListBooking)
	booking.Post("/create", handler.CreateBooking)
	booking.Delete("/delete/:id", handler.DeleteBooking)

	// booking seat
	booking_seat := api.Group("/booking/seat")
	booking_seat.Get("/", handler.ListBookingSeat)
	booking_seat.Post("/create", handler.CreateBookingSeat)
	booking_seat.Delete("/delete/:id", handler.DeleteBookingSeat)

	// payment
	payment := api.Group("/payment")
	payment.Post("/", handler.ListPayment)
	payment.Post("/create", handler.CreatePayment)
}
