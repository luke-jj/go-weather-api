package startup

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/luke-jj/go-weather-api/internal/models"
)

func Routes(config *models.Config, router *chi.Mux) {
}

func LogRoutes(router *chi.Mux) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}
}
