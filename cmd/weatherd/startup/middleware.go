package startup

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	m "github.com/luke-jj/go-weather-api/cmd/weatherd/middleware"
	c "github.com/luke-jj/go-weather-api/internal/config"
	d "github.com/luke-jj/go-weather-api/internal/database"
)

func Middleware(config *c.Config, db *d.Database, r *chi.Mux) {
	r.Use(
		// TODO: set config as context
		m.GetCORSHandler(),
		m.SetConfigContext(config),
		m.SetDatabaseContext(db),
		middleware.SetHeader("Content-Type", "application/json; charset=utf-8"),
		middleware.Timeout(20*time.Second),
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Compress(5),
		middleware.Recoverer,
		middleware.URLFormat,
	)
}
