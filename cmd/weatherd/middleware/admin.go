package middleware

import "fmt"
import "net/http"

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Checking admin status.")
		next.ServeHTTP(w, r)
	})
}
