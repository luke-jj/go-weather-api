package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/luke-jj/go-weather-api/cmd/weatherd/middleware"
	"github.com/luke-jj/go-weather-api/cmd/weatherd/models"
	c "github.com/luke-jj/go-weather-api/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		var forecast models.Forecast
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		forecastsCollection := config.Db.Collection("forecasts")

		// TODO: Validate request body

		json.NewDecoder(r.Body).Decode(&forecast)
		result, err := forecastsCollection.InsertOne(ctx, forecast)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
			return
		}
		// TODO: insert type assertion
		// forecast.ID = result.InsertedID
		fmt.Println(result)
		render.JSON(w, r, forecast)
	}
}

func getForecasts(config *c.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var forecasts []models.Forecast
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		forecastsCollection := config.Db.Collection("forecasts")

		cursor, err := forecastsCollection.Find(ctx, bson.M{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
			return
		}
		if err = cursor.All(ctx, &forecasts); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
			return
		}
		// returns null instead of an empty array if no data in db
		render.JSON(w, r, forecasts)
	}
}

func getForecastById(config *c.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var forecast models.Forecast
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		forecastsCollection := config.Db.Collection("forecasts")
		id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
		if err != nil {
			// TODO: write a 400 - bad request instead
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
			return
		}

		result := forecastsCollection.FindOne(ctx, models.Forecast{ID: id})
		// TODO: return 404 if no result found
		fmt.Println(result)

		err = result.Decode(&forecast)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
			return
		}
		render.JSON(w, r, forecast)
	}
}

func updateForecast(config *c.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var forecast models.Forecast
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		forecastsCollection := config.Db.Collection("forecasts")
		id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
		if err != nil {
			// TODO: write a 400 - bad request instead
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
			return
		}

		// TODO: Validate request body

		result := forecastsCollection.FindOne(ctx, models.Forecast{ID: id})
		// TODO: return 404 if no result found
		fmt.Println(result)

		res, err := forecastsCollection.UpdateOne(
			ctx,
			// update by id, can use other properties here.
			bson.M{"_id": id},
			bson.D{
				{"$set", bson.D{{"author", "Nicolas Raboi"}}},
			},
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
			return
		}
	}
}

func deleteForecast(config *c.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var forecast models.Forecast
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		forecastsCollection := config.Db.Collection("forecasts")
		id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
		if err != nil {
			// TODO: write a 400 - bad request instead
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
			return
		}

		result := forecastsCollection.FindOne(ctx, models.Forecast{ID: id})
		// TODO: return 404 if no result found
		fmt.Println(result)
	}
}
