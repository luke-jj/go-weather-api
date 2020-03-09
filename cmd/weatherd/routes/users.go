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
	"github.com/luke-jj/go-weather-api/cmd/weatherd/middleware"
	"github.com/luke-jj/go-weather-api/cmd/weatherd/models"
	d "github.com/luke-jj/go-weather-api/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Users() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", createUser)
	r.With(middleware.Admin).Get("/", getUsers)
	r.Route("/{id}", func(r chi.Router) {
		r.With(middleware.Admin).Get("/", getUserById)
		r.With(middleware.Admin).Put("/", updateUser)
		r.With(middleware.Admin).Delete("/", deleteUser)
	})

	return r
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	validate := validator.New()
	err := validate.Struct(user)
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
	err = coll.FindOne(ctx, models.User{Email: user.Email}).Decode(&user)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{ "message": "` + "Email already in use." + `"}`))
		return
	}
	if err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	user.Password = string(hash)

	result, err := coll.InsertOne(ctx, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	user.ID = id
	render.JSON(w, r, user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db, ok := r.Context().Value("db").(*d.Database)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	coll := db.Db.Collection("users")
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	if err = cursor.All(ctx, &users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	for i := range users {
		users[i].Password = ""
	}
	// renders 'null' instead of an empty array '[]' if no data in db.
	render.JSON(w, r, users)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	var user models.User
	db, ok := r.Context().Value("db").(*d.Database)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	coll := db.Db.Collection("users")
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{ "message": "` + "Not a valid id." + `"}`))
		return
	}
	err = coll.FindOne(ctx, models.User{ID: id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{ "message": "User with given id not found."}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	user.Password = ""
	render.JSON(w, r, user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var replacedUser models.User
	json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Input validation failed. Field '%v' must be of type '%v' and satisfy the condition: '%v %v'", err.Field(), err.Kind().String(), err.Tag(), err.Param())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{ "message": "` + message + `"}`))
			return
		}
	}
	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
			return
		}
		user.Password = string(hash)
	}
	db, ok := r.Context().Value("db").(*d.Database)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	coll := db.Db.Collection("users")
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{ "message": "` + "Not a valid id." + `"}`))
		return
	}
	filter := bson.M{"_id": id}
	user.ID = id
	err = coll.FindOneAndReplace(ctx, filter, user).Decode(&replacedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{ "message": "User with given id not found."}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	user.Password = ""
	render.JSON(w, r, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var deletedUser models.User
	db, ok := r.Context().Value("db").(*d.Database)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	coll := db.Db.Collection("users")
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{ "message": "` + "Not a valid id." + `"}`))
		return
	}
	err = coll.FindOneAndDelete(ctx, bson.D{{"_id", id}}).Decode(&deletedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{ "message": "User with given id not found."}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + http.StatusText(500) + `"}`))
		return
	}
	deletedUser.Password = ""
	render.JSON(w, r, deletedUser)
}
