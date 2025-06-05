package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	userToken, err := h.domain.GoogleOAuthLogin(ctx, code)
	if err != nil {
		if err.Error() == "not found" {
			redirectURL := fmt.Sprintf(
				"http://local.fidelity.com:%s/login?error=user_not_found",
				h.ClientPort,
			)
			c.Redirect(http.StatusTemporaryRedirect, redirectURL)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	redirectURL := fmt.Sprintf(
		"http://local.fidelity.com:%s/login?token=%s",
		h.ClientPort,
		userToken,
	)
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	return
}

func (h *AuthHandler) IsFirstLogin(c *gin.Context) {
	ctx := c.Request.Context()
	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing email"})
		return
	}

	c.JSON(http.StatusOK, h.domain.IsFirstLogin(ctx, email))
}
