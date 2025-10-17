package routes

import (
	"cinema/internal/app/handlers"
	"cinema/config"
	"net/http"
)

func Setup(mux *http.ServeMux) {
	r := &config.Router{mux}
	r.Get("/test", handlers.GetTest)
	r.Get("/tulit", handlers.GetTulit)
	r.Get("/data", handlers.GetData)
}
