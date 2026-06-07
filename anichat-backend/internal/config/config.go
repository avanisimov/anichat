package config

import (
	"errors"
	"os"
)

type Config struct {
	Port               string
	AppURL             string
	DatabaseURL        string
	ResendAPIKey       string
	JWTSecret          string
}

func Load() (*Config, error) {
	cfg := &Config{
		Port:               lookupEnv("PORT", "8080"),
		AppURL:             os.Getenv("APP_URL"),
		DatabaseURL:        os.Getenv("DATABASE_URL"),
		ResendAPIKey:       os.Getenv("RESEND_API_KEY"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
	}

	if cfg.AppURL == "" || cfg.DatabaseURL == "" || cfg.ResendAPIKey == "" || cfg.JWTSecret == "" {
		return nil, errors.New("one or more required environment variables are missing")
	}

	return cfg, nil
}

func lookupEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
