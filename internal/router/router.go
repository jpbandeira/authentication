package router

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/jp/authentication/internal/handler"

	"github.com/gin-gonic/gin"
)

func Setup(ctx context.Context, handler *handler.AuthHandler) *gin.Engine {
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.Use(cors.Default())

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
	router.GET("/auth/google/callback", handler.GoogleCallbackHandler)

	return router
}
