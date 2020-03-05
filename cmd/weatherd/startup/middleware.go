package startup

import (
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/luke-jj/go-weather-api/cmd/weatherd/middleware"
	"github.com/luke-jj/go-weather-api/internal/models"
)

func Middleware(config *models.Config, router *chi.Mux) {
	router.Use(
		middleware.GetCORSHandler(),
		render.SetContentType(render.ContentTypeJSON),
		chiMiddleware.Logger,
		chiMiddleware.Compress(5),
		chiMiddleware.RedirectSlashes,
		chiMiddleware.Recoverer,
	)
}
