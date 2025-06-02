package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jp/authentication/internal/service"
	sModel "github.com/jp/authentication/internal/service/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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

	_ = godotenv.Load()

	var googleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://local.fidelity.com:8083/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/calendar.readonly",
			"https://www.googleapis.com/auth/userinfo.email",
			"openid",
		},
		Endpoint: google.Endpoint,
	}

	token, err := googleOAuthConfig.Exchange(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token exchange failed", "details": err.Error()})
		return
	}

	client := googleOAuthConfig.Client(ctx, token)
	email, err := service.FetchUserEmail(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch email", "details": err.Error()})
		return
	}

	user, err := h.repo.FindByEmail(ctx, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user registers", "details": err.Error()})
		return
	}

	userEmail, err := h.repo.SaveGoogleToken(ctx, sModel.GoogleToken{
		ID:           user.ID,
		UserEmail:    user.Email,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user google token", "details": err.Error()})
	}

	redirectURL := fmt.Sprintf("http://local.fidelity.com:8082/client?email=%s", url.QueryEscape(userEmail))
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
