package middleware

import (
	"context"
	"net/http"

	d "github.com/luke-jj/go-weather-api/internal/database"
)

func SetDatabaseContext(db *d.Database) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
