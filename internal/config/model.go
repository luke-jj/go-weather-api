package config

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
