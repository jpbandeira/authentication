package config

import (
	"log"
	"os"
)

func Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
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
