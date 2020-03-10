package middleware

import (
	"github.com/luke-jj/go-weather-api/pkg/models"
	"net/http"
)

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(*models.Claims)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
			return
		}
		if !user.IsAdmin {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{ "message": "` + http.StatusText(403) + `"}`))
			return
		}
		next.ServeHTTP(w, r)
	})
}
