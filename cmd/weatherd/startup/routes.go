package startup

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/luke-jj/go-weather-api/cmd/weatherd/routes"
)

func Routes(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/forecasts", routes.Forecasts())
		r.Mount("/users", routes.Users())
		r.Mount("/auth", routes.Auth())
		r.Mount("/time", routes.Times())
		r.Mount("/weather", routes.Weather())
	})
}

func LogRoutes(r *chi.Mux) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}
}
