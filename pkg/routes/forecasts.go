package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	d "github.com/luke-jj/go-weather-api/internal/database"
	"github.com/luke-jj/go-weather-api/pkg/middleware"
	"github.com/luke-jj/go-weather-api/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Forecasts() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getForecasts)
	r.With(middleware.Auth, middleware.Admin).Post("/", createForecast)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", getForecastById)
		r.With(middleware.Auth, middleware.Admin).Put("/", updateForecast)
		r.With(middleware.Auth, middleware.Admin).Delete("/", deleteForecast)
	})

	return r
}

func createForecast(w http.ResponseWriter, r *http.Request) {
	var forecast models.Forecast
	json.NewDecoder(r.Body).Decode(&forecast)
	defer r.Body.Close()
	validate := validator.New()
	err := validate.Struct(forecast)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Input validation failed. Field '%v' must be of type '%v' and satisfy the condition: '%v %v'", err.Field(), err.Kind().String(), err.Tag(), err.Param())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{ "message": "` + message + `"}`))
			return
		}
	}

	db, ok := r.Context().Value("db").(*d.Database)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	coll := db.Db.Collection("forecasts")
	result, err := coll.InsertOne(ctx, forecast)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	forecast.ID = id
	render.JSON(w, r, forecast)
}

func getForecasts(w http.ResponseWriter, r *http.Request) {
	var forecasts []models.Forecast
	db, ok := r.Context().Value("db").(*d.Database)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	coll := db.Db.Collection("forecasts")
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	if err = cursor.All(ctx, &forecasts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	// renders 'null' instead of an empty array '[]' if no data in db.
	render.JSON(w, r, forecasts)
}

func getForecastById(w http.ResponseWriter, r *http.Request) {
	var forecast models.Forecast
	db, ok := r.Context().Value("db").(*d.Database)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	coll := db.Db.Collection("forecasts")
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{ "message": "` + "Not a valid id." + `"}`))
		return
	}
	err = coll.FindOne(ctx, models.Forecast{ID: id}).Decode(&forecast)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{ "message": "Forecast with given id not found."}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	render.JSON(w, r, forecast)
}

func updateForecast(w http.ResponseWriter, r *http.Request) {
	var forecast models.Forecast
	var replacedForecast models.Forecast
	json.NewDecoder(r.Body).Decode(&forecast)
	defer r.Body.Close()
	validate := validator.New()
	err := validate.Struct(forecast)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Input validation failed. Field '%v' must be of type '%v' and satisfy the condition: '%v %v'", err.Field(), err.Kind().String(), err.Tag(), err.Param())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{ "message": "` + message + `"}`))
			return
		}
	}
	db, ok := r.Context().Value("db").(*d.Database)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	coll := db.Db.Collection("forecasts")
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{ "message": "` + "Not a valid id." + `"}`))
		return
	}
	filter := bson.M{"_id": id}
	forecast.ID = id
	err = coll.FindOneAndReplace(ctx, filter, forecast).Decode(&replacedForecast)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{ "message": "Forecast with given id not found."}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	render.JSON(w, r, forecast)
}

func deleteForecast(w http.ResponseWriter, r *http.Request) {
	var deletedForecast models.Forecast
	db, ok := r.Context().Value("db").(*d.Database)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	coll := db.Db.Collection("forecasts")
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{ "message": "` + "Not a valid id." + `"}`))
		return
	}
	err = coll.FindOneAndDelete(ctx, bson.D{{"_id", id}}).Decode(&deletedForecast)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{ "message": "Forecast with given id not found."}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	render.JSON(w, r, deletedForecast)
}
