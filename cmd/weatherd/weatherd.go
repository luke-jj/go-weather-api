package weatherd

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/luke-jj/go-weather-api/cmd/weatherd/startup"
	"github.com/luke-jj/go-weather-api/internal/config"
	"github.com/luke-jj/go-weather-api/internal/database"
)

func Startup() {
	router := chi.NewRouter()
	config := config.Read()
	database.Init(config)
	defer config.Client.Disconnect(config.Ctx)
	startup.Middleware(config, router)
	startup.Routes(config, router)
	startup.LogRoutes(router)
	log.Printf("Listening on port %v...\n", config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, router))
}
