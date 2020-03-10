package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	c "github.com/luke-jj/go-weather-api/internal/config"
	"github.com/luke-jj/go-weather-api/internal/database"
	"github.com/luke-jj/go-weather-api/pkg/startup"
)

func main() {
	router := chi.NewRouter()
	config := c.Read()
	db := database.Init(config)
	defer db.Client.Disconnect(db.Ctx)
	startup.Middleware(config, db, router)
	startup.Routes(router)
	startup.LogRoutes(router)
	log.Printf("Listening on port %v...\n", config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, router))
}
