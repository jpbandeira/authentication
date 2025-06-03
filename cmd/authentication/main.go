package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/jp/authentication/internal/config"
	"github.com/jp/authentication/internal/database"
	"github.com/jp/authentication/internal/domain"
	"github.com/jp/authentication/internal/handler"
	"github.com/jp/authentication/internal/repository"
	"github.com/jp/authentication/internal/router"
)

func main() {
	_ = godotenv.Load()
	ctx := context.Background()

	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	repo := repository.NewGormRepository(db)
	domain := domain.NewDomain(repo, config.JWTSecret())
	authHandler := handler.NewAuthHandler(domain, config.ClientPort())

	r := router.Setup(ctx, authHandler)
	r.Run(":" + config.HostPort())
}
