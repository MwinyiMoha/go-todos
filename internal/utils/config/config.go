package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	AppTimeout     int64
	Port           string
	DBName         string
	CollectionName string
	DBHost         string
	DBUser         string
	DBPassword     string
	DBPort         string
	AuthDB         string
}

func MustGetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return defaultValue
}

func New() *Config {
	appTimeout := MustGetEnv("APP_TIMEOUT", "10")
	timeout, err := strconv.ParseInt(appTimeout, 0, 64)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &Config{
		AppTimeout:     timeout,
		Port:           MustGetEnv("PORT", "8080"),
		DBName:         MustGetEnv("DB_NAME", "todos"),
		CollectionName: MustGetEnv("COLLECTION", "todos"),
		DBHost:         MustGetEnv("DB_HOST", "localhost"),
		DBUser:         MustGetEnv("DB_USER", "user"),
		DBPassword:     MustGetEnv("DB_PASSWORD", "password"),
		DBPort:         MustGetEnv("DB_PORT", "27017"),
		AuthDB:         MustGetEnv("AUTH_DB", "admin"),
	}
}
