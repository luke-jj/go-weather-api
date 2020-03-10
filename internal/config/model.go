package config

type Config struct {
	ENVIRONMENT_MODE string `validate:"required"`
	CONFIG_NAME      string
	PORT             string `validate:"required"`
	WEATHER_URI      string `validate:"required"`
	WEATHER_KEY      string `validate:"required"`
	TIME_URI         string `validate:"required"`
	MONGO_URI        string `validate:"required"`
	MONGO_DBNAME     string `validate:"required"`
	JWTPRIVATEKEY    string `validate:"required"`
	LOG_ENABLED      bool
}
