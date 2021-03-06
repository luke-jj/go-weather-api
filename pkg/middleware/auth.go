package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	c "github.com/luke-jj/go-weather-api/internal/config"
	"github.com/luke-jj/go-weather-api/pkg/models"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["X-Auth-Token"] == nil {
			w.WriteHeader(401)
			w.Write([]byte(`{ "message": "Access denied. No token provided."}`))
			return
		}
		config, ok := r.Context().Value("config").(*c.Config)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
			return
		}
		token, err := jwt.ParseWithClaims(r.Header["X-Auth-Token"][0], &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Token uses different signing method.")
			}
			return []byte(config.JWTPRIVATEKEY), nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(401)
			w.Write([]byte(`{ "message": "Access denied. Invalid token."}`))
			return
		}
		ctx := context.WithValue(r.Context(), "user", token.Claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
