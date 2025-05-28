package main

import (
	"context"
	"log"

	"github.com/jp/authentication/internal/config"
	"github.com/jp/authentication/internal/database"
	"github.com/jp/authentication/internal/router"
)

func main() {
	ctx := context.Background()

	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	r := router.Setup(ctx, db)
	r.Run(":" + config.Port())
}
