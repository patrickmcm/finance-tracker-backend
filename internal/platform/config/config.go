package config

import (
	"log"
	"os"
)

type Config struct {
	DBUser string
	DBPass string
}

func LoadConfig() *Config {
	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("DB_USER ENV VAR REQUIRED")
	}
	password := os.Getenv("DB_PASS")
	if password == "" {
		log.Fatal("DB_USER ENV VAR REQUIRED")
	}

	return &Config{
		DBUser: user,
		DBPass: password,
	}
}
