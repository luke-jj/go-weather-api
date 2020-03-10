package middleware

import (
	"context"
	"net/http"

	c "github.com/luke-jj/go-weather-api/internal/config"
)

func SetConfigContext(config *c.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "config", config)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
