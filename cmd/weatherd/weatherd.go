package weatherd

import (
	"fmt"
	"log"
	// "net/http"

	"github.com/luke-jj/go-weather-api/internal/config"
)

func Startup() {
	config, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening on port %v...\n", config.PORT)
	// log.Fatal(http.ListenAndServe(":"+config.PORT, r)
}
