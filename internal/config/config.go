package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	ENVIRONMENT_MODE string
	CONFIG_NAME      string
	PORT             string
	WEATHER_URI      string
	WEATHER_KEY      string
	TIME_URI         string
	MONGO_URI        string
	MONGO_DBNAME     string
	JWTPRIVATEKEY    string
	LOG_ENABLED      bool
}

func New() (*Config, error) {
	var err error = nil
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
		err = errors.New("FATAL ERROR: PORT must be set.")
	}
	if envMode == "" {
		err = errors.New("FATAL ERROR: ENVIRONMENT_MODE must be set.")
	}
	if weatherURI == "" {
		err = errors.New("FATAL ERROR: API_WEATHER_URI must be set.")
	}
	if weatherKey == "" {
		err = errors.New("FATAL ERROR: API_WEATHER_KEY must be set.")
	}
	if timeURI == "" {
		err = errors.New("FATAL ERROR: API_TIME_URI must be set.")
	}
	if mongoURI == "" {
		err = errors.New("FATAL ERROR: API_MONGO_URI must be set.")
	}
	if dbName == "" {
		err = errors.New("FATAL ERROR: API_MONGO_DBNAME must be set.")
	}
	if jwtKey == "" {
		err = errors.New("FATAL ERROR: API_JWTPRIVATEKEY must be set.")
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

	return &config, err
}
