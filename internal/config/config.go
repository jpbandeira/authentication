package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Port() string {
	_ = godotenv.Load()

	port := os.Getenv("AUTH_PORT")
	if port == "" {
		port = "8083"
	}
	return port
}

func JWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET não definido")
	}
	return secret
}
