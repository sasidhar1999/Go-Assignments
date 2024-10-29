package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDSN string
	KafkaBroker string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// here values are taken from .env file
	return Config{
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
	}
}
