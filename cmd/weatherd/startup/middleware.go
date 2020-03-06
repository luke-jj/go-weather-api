package startup

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	m "github.com/luke-jj/go-weather-api/cmd/weatherd/middleware"
	c "github.com/luke-jj/go-weather-api/internal/config"
)

func Middleware(config *c.Config, r *chi.Mux) {
	r.Use(
		m.SetContentTypeJSON,
		m.GetCORSHandler(),
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Compress(5),
		middleware.Recoverer,
	)
}
