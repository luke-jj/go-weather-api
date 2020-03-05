package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/luke-jj/go-weather-api/cmd/weatherd/middleware"
	"github.com/luke-jj/go-weather-api/cmd/weatherd/models"
	c "github.com/luke-jj/go-weather-api/internal/config"
)

func Forecasts(config *c.Config) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", getForecasts(config))
	r.Get("/{id}", getForecastById(config))
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.Admin)
		r.Post("/", createForecast(config))
		r.Put("/{id}", updateForecast(config))
		r.Delete("/{id}", deleteForecast(config))
	})

	return r
}

func createForecast(config *c.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := make(map[string]string)
		response["message"] = "Created Forecast successfully"
		render.JSON(w, r, response)
	}
}

func getForecasts(config *c.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		forecasts := []models.Forecast{
			{
				// ID:    "1902831049",
				Title: "Hello world",
				Text:  "Heloo world from planet earth",
			},
			{
				// ID:    "298458949",
				Title: "Hello world",
				Text:  "Heloo world from planet earth",
			},
			{
				// ID:    "9038450385",
				Title: "Hello world",
				Text:  "Heloo world from planet earth",
			},
		}
		render.JSON(w, r, forecasts)
	}
}

func getForecastById(config *c.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Getting forecast with id %v...", chi.URLParam(r, "id"))))
	}
}

func updateForecast(config *c.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Updating forecast with id %v...", chi.URLParam(r, "id"))))
	}
}

func deleteForecast(config *c.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Deleting forecast with id %v...", chi.URLParam(r, "id"))))
	}
}
