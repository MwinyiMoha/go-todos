package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	AppTimeout int64
	Port       string
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
		AppTimeout: timeout,
		Port:       MustGetEnv("PORT", "8080"),
	}
}
