package routes

import (
	"cinema/config"
	"cinema/internal/app/handlers"
	"net/http"
)

func Setup(mux *http.ServeMux) {
	r := &config.Router{mux}
	r.Get("/test", handlers.GetTest)
	r.Get("/tulit", handlers.GetTulit)
	r.Get("/data", handlers.GetData)
}
