package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"github.com/luke-jj/go-weather-api/cmd/weatherd/models"
	c "github.com/luke-jj/go-weather-api/internal/config"
	d "github.com/luke-jj/go-weather-api/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Auth() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var auth models.UserAuthValidation
		var user models.User
		json.NewDecoder(r.Body).Decode(&auth)
		defer r.Body.Close()
		validate := validator.New()
		err := validate.Struct(auth)
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
		coll := db.Db.Collection("users")
		err = coll.FindOne(ctx, models.User{Email: auth.Email}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{ "message": "Invalid username or password."}`))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{ "message": "Invalid username or password."}`))
			return
		}

		config, ok := r.Context().Value("config").(*c.Config)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
			return
		}
		token, err := user.GenerateAuthToken(config.JWTPRIVATEKEY)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
			return
		}
		w.Write([]byte(`{ "token": "` + token + `"}`))
	})

	return r
}
