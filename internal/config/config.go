package config

import (
	"fmt"
	"os"

	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	TokenAuth   *jwtauth.JWTAuth
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}
	tokenAuth := jwtauth.New("HS256", []byte(jwtSecret), nil)
	cfg := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		TokenAuth:   tokenAuth,
	}

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	return cfg, nil
}
