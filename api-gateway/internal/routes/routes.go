package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yokeTH/gofiber-scalar/scalar/v2"

	"api/config"
	"api/internal/app/handler"
	"api/internal/middleware"
)

func Setup(app *fiber.App) {
	app.Get("/", handler.Home)

	api := app.Group("/api/v1")
	api.Use(config.AllowDoc)
	api.Get("/docs/*", scalar.New(config.APIDocs()))

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
	genre := api.Group("/genre", middleware.Permission("handle genre"))
	api.Get("/genre/", handler.ListGenre)
	api.Get("/genre/:id", handler.GetGenre)
	genre.Post("/create", handler.CreateGenre)
	genre.Put("/update/:id", handler.UpdateGenre)
	genre.Delete("/delete/:id", handler.DeleteGenre)

	// film
	api.Get("/film/", handler.ListFilm)
	api.Get("/film/:id", handler.GetFilm)
	api.Post("/film/search", handler.SearchFilm)
	api.Get("/film/genre/:id", handler.SearchFilmGenre)
	film := api.Group("/film", middleware.Permission("handle film"))
	film.Post("/create", handler.CreateFilm)
	film.Put("/update/:id", handler.UpdateFilm)
	film.Delete("/delete/:id", handler.DeleteFilm)

	// cinema
	api.Get("/cinema/", handler.ListCinema)
	api.Get("/cinema/schedule/:id", handler.ListCinemaSchedule)
	api.Get("/cinema/:id", handler.GetCinema)
	cinema := api.Group("/cinema", middleware.Permission("handle cinema"))
	cinema.Post("/create", handler.CreateCinema)
	cinema.Put("/update/:id", handler.UpdateCinema)
	cinema.Delete("/delete/:id", handler.DeleteCinema)

	// screen type
	api.Get("/screen/type/", handler.ListScreenType)
	api.Get("/screen/type/:id", handler.GetScreenType)
	screen_type := api.Group("/screen/type", middleware.Permission("handle screen type"))
	screen_type.Post("/create", handler.CreateScreenType)
	screen_type.Put("/update/:id", handler.UpdateScreenType)
	screen_type.Delete("/delete/:id", handler.DeleteScreenType)

	// screen
	api.Get("/screen", handler.ListScreen)
	api.Get("/screen/:id", handler.GetScreen)
	screen := api.Group("/screen", middleware.Permission("handle screen"))
	screen.Post("/create", handler.CreateScreen)
	screen.Put("/update/:id", handler.UpdateScreen)
	screen.Delete("/delete/:id", handler.DeleteScreen)

	// seat
	api.Get("/seat/", handler.ListSeat)
	api.Get("/seat/:id", handler.GetSeat)
	seat := api.Group("/seat", middleware.Permission("handle seat"))
	seat.Post("/create", handler.CreateSeat)
	seat.Put("/update/:id", handler.UpdateSeat)
	seat.Delete("/delete/:id", handler.DeleteSeat)

	// showtime
	api.Get("/showtime/", handler.ListShowTime)
	api.Get("/showtime/:id", handler.GetShowTime)
	showtime := api.Group("/showtime", middleware.Permission("handle showtime"))
	showtime.Post("/create", handler.CreateShowTime)
	showtime.Put("/update/:id", handler.UpdateShowTime)
	showtime.Delete("/delete/:id", handler.DeleteShowTime)

	// booking
	booking := api.Group("/booking")
	booking.Get("/", handler.ListBooking)
	booking.Post("/create", handler.CreateBooking)
	booking.Delete("/delete/:id", handler.DeleteBooking)

	// booking seat
	booking_seat := api.Group("/booking/seat")
	booking_seat.Post("/", handler.ListBookingSeat)
	booking_seat.Post("/create", handler.CreateBookingSeat)
	booking_seat.Delete("/delete/:id", handler.DeleteBookingSeat)

	// payment
	payment := api.Group("/payment")
	payment.Post("/", handler.ListPayment)
	payment.Post("/create", handler.CreatePayment)
}
