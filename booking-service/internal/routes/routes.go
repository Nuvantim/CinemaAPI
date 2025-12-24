package routes

import (
	"booking/config"
	"booking/internal/app/handler"
	"net/http"
)

func Setup(mux *http.ServeMux) {
	r := &config.Router{Mux: mux}
	r.Get("/", handler.GetTest)

	// booking
	r.Post("/bookings", handler.ListBooking)
	r.Get("/booking/", handler.GetBooking)
	r.Post("/booking/create", handler.CreateBooking)
	r.Delete("/booking/delete/", handler.DeleteBooking)

	// booking seat
	r.Post("/booking/seats", handler.ListBookingSeat)
	r.Post("/booking/seat/create", handler.CreateBookingSeat)
	r.Delete("/booking/seat/delete/", handler.DeleteBookingSeat)

	// payment
	r.Post("/payment/create", handler.CreatePayment)
	r.Post("/payment/", handler.ListPayment)
}
