package config

import (
	"log"
	"os"
)

func HostPort() string {
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
func ClientPort() string {
	clientPort := os.Getenv("CLIENT_PORT")
	if clientPort == "" {
		log.Fatal("CLIENT_PORT não definido")
	}
	return clientPort
}
