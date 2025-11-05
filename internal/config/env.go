package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppServer  string
	AppPort string
	PostGresConnString string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	env := os.Getenv("APP_ENV")
	port := os.Getenv("APP_PORT")
	dns := os.Getenv("POSTGRES_DSN")
	server := os.Getenv("APP_SERVER")

	if env == "" {
		env = "development"
	}
	if port == "" {
		port = "8080"
	}
	return &Config{
		AppEnv:  env,
		AppServer:  server+":"+port,
		AppPort: ":"+port,
		PostGresConnString: dns,
	}
}
