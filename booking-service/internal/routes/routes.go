package routes

import (
	"booking/config"
	"booking/internal/app/handlers"
	"net/http"
)

func Setup(mux *http.ServeMux) {
	r := &config.Router{mux}
	r.Get("/", handler.GetTest)

	// booking
	r.Post("/bookings", handler.ListBooking)
	r.Post("/booking/create", handler.CreateBooking)
	r.Delete("/booking/delete/", handler.DeleteBooking)

	// booking seat
	r.Get("/booking/seats", handler.ListBookingSeat)
	r.Post("/booking/seat/create", handler.CreateBookingSeat)
	r.Delete("/booking/seat/delete/", handler.DeleteBookingSeat)

	// payment
	r.Post("/payment/create", handler.CreatePayment)
	r.Get("/payment/", handler.ListPayment)
}
