package router

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/jp/authentication/internal/handler"

	"github.com/gin-gonic/gin"
)

const (
	basePath string = "/authentication"
)

func Setup(ctx context.Context, handler *handler.AuthHandler) *gin.Engine {
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.Use(cors.Default())

	router.POST(basePath+"/user/register", handler.Register)
	router.GET(basePath+"/user/:email", handler.FindByEmail)
	router.POST(basePath+"/login", handler.Login)
	router.GET(basePath+"/auth/google/callback", handler.GoogleCallbackHandler)
	router.GET(basePath+"/first-login/:email", handler.IsFirstLogin)

	return router
}
