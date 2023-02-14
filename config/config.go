package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port        int
	DatabaseURL string
}

func Load() *Config {
	portStr := os.Getenv("APP_PORT")
	if portStr == "" {
		portStr = "8080"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid port number: %s", portStr))
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		panic("Missing DATABASE_URL environment variable")
	}

	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
	}
}
