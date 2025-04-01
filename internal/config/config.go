package config

import (
	"log"

	"github.com/alphatechnolog/purplish-items/pkg/helpers"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
}

func LoadConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		log.Printf("Error loading env file: %v", err)
	}

	cfg := &Config{
		DatabaseURL: helpers.GetEnv("DATABASE_URL", "database.db"),
		ServerPort:  helpers.GetEnv("SERVER_PORT", "8001"),
	}

	return cfg, nil
}
