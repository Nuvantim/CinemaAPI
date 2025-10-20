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
	r.Get("/bookings", handler.ListBooking)
	r.Post("/booking/create", handler.CreateBooking)
	r.Delete("/booking/delete/", handler.DeleteBooking)

	// booking seat
	// r.Get("/booking/seats", handler.ListBookingSeat)
	// r.Post("/booking/create", handler.CreateBookingSeat)

	// payment
	// r.Post("/payment/create", handler.CreatePayment)
	// r.Get("/payment/", handler.ListPayment)
	// r.Get("/payment/report", handler.ReportPayment)
}
