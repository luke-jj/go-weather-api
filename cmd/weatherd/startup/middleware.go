package startup

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	mw "github.com/luke-jj/go-weather-api/cmd/weatherd/middleware"
	"github.com/luke-jj/go-weather-api/internal/models"
)

func Middleware(config *models.Config, router *chi.Mux) {
	router.Use(
		mw.GetCORSHandler(),
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Compress(5),
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)
}
