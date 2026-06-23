package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	Production  = "production"
	Development = "development"
)

var (
	DBHost              string
	DBPort              string
	DBUser              string
	DBPassword          string
	DBName              string
	DBSSLMode           string
	ServerPort          string
	TokenSecret         string
	TokenExpirationMins int
	APPEnv              string
)

func Load() error {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: No se encontró el archivo .env, usando variables del sistema")
	}

	// Si APP_ENV es production, cargar .env.prod (pisa valores de .env)
	if os.Getenv("APP_ENV") == Production || os.Getenv("APP_ENV") == "" {
		_ = godotenv.Overload(".env.prod")
	}

	DBHost = getEnv("DB_HOST", "localhost")
	DBPort = getEnv("DB_PORT", "5432")
	DBUser = getEnv("DB_USER", "postgres")
	DBPassword = getEnv("DB_PASSWORD", "")
	DBName = getEnv("DB_NAME", "idea")
	DBSSLMode = getEnv("DB_SSLMODE", "disable")
	ServerPort = getEnv("PORT", "8080")
	APPEnv = getEnv("APP_ENV", "development")

	return nil
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if val := os.Getenv(key); val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return defaultValue
}
