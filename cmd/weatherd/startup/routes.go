package startup

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/luke-jj/go-weather-api/cmd/weatherd/routes"
	c "github.com/luke-jj/go-weather-api/internal/config"
)

func Routes(config *c.Config, r *chi.Mux) {
	r.Mount("/api/v1/forecasts", routes.Forecasts(config))
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
