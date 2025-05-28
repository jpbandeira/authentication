package router

import (
	"context"

	"github.com/jp/authentication/internal/handler"
	"github.com/jp/authentication/internal/repository"
	"github.com/jp/authentication/internal/service"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Setup(ctx context.Context, db *gorm.DB) *gin.Engine {
	r := gin.Default()
	repo := repository.NewUserRepository(db)
	authService := service.NewAuthService()
	authHandler := handler.NewAuthHandler(repo, authService)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	return r
}
