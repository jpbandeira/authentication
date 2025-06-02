package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jp/authentication/internal/oauth"
	"github.com/jp/authentication/internal/service"
	sModel "github.com/jp/authentication/internal/service/model"
)

type GoogleAuth struct {
	Email string
}

func (h *AuthHandler) GoogleCallbackHandler(c *gin.Context) {
	ctx := c.Request.Context()
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}

	token, err := oauth.GoogleOAuthConfig.Exchange(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token exchange failed", "details": err.Error()})
		return
	}

	client := oauth.GoogleOAuthConfig.Client(ctx, token)
	email, err := service.FetchUserEmail(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch email", "details": err.Error()})
		return
	}

	userEmail, err := h.repo.SaveGoogleToken(ctx, sModel.GoogleToken{
		UserEmail:    email,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user google token", "details": err.Error()})
	}

	c.JSON(http.StatusOK, GoogleAuth{Email: userEmail})
}
