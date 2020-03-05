package config

import (
	"errors"
	"log"
	"os"
	"strconv"
)

func Read() *Config {
	port := os.Getenv("PORT")
	envMode := os.Getenv("ENVIRONMENT_MODE")
	name := os.Getenv("API_CONFIG_NAME")
	weatherURI := os.Getenv("API_WEATHER_URI")
	weatherKey := os.Getenv("API_WEATHER_KEY")
	timeURI := os.Getenv("API_TIME_URI")
	mongoURI := os.Getenv("API_MONGO_URI")
	dbName := os.Getenv("API_MONGO_DBNAME")
	jwtKey := os.Getenv("API_JWTPRIVATEKEY")
	logEnabled, _ := strconv.ParseBool(os.Getenv("API_LOG_ENABLED"))

	if port == "" {
		log.Fatal(errors.New("FATAL ERROR: PORT must be set."))
	}
	if envMode == "" {
		log.Fatal(errors.New("FATAL ERROR: ENVIRONMENT_MODE must be set."))
	}
	if weatherURI == "" {
		log.Fatal(errors.New("FATAL ERROR: API_WEATHER_URI must be set."))
	}
	if weatherKey == "" {
		log.Fatal(errors.New("FATAL ERROR: API_WEATHER_KEY must be set."))
	}
	if timeURI == "" {
		log.Fatal(errors.New("FATAL ERROR: API_TIME_URI must be set."))
	}
	if mongoURI == "" {
		log.Fatal(errors.New("FATAL ERROR: API_MONGO_URI must be set."))
	}
	if dbName == "" {
		log.Fatal(errors.New("FATAL ERROR: API_MONGO_DBNAME must be set."))
	}
	if jwtKey == "" {
		log.Fatal(errors.New("FATAL ERROR: API_JWTPRIVATEKEY must be set."))
	}

	config := Config{
		ENVIRONMENT_MODE: envMode,
		PORT:             port,
		CONFIG_NAME:      name,
		WEATHER_URI:      weatherURI,
		WEATHER_KEY:      weatherKey,
		TIME_URI:         timeURI,
		MONGO_URI:        mongoURI,
		MONGO_DBNAME:     dbName,
		JWTPRIVATEKEY:    jwtKey,
		LOG_ENABLED:      logEnabled,
	}

	return &config
}
