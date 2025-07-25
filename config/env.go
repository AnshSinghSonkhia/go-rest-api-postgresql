package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		DBHost:     getENV("DB_HOST", "127.0.0.1"),
		DBPort:     getENV("DB_PORT", "5432"),
		DBUser:     getENV("DB_USER", "user"),
		DBPassword: getENV("DB_PASSWORD", "password"),
		DBName:     getENV("DB_NAME", "postgres"),
	}
}

func getENV(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
