package router

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/jp/authentication/internal/handler"
	"github.com/jp/authentication/internal/repository"
	"github.com/jp/authentication/internal/service"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Setup(ctx context.Context, db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.Use(cors.Default())
	repo := repository.NewRepository(db)
	authService := service.NewAuthService()
	authHandler := handler.NewAuthHandler(repo, authService)

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
	router.POST("/auth/google/callback", authHandler.GoogleCallbackHandler)

	return router
}
