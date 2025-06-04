package handler

import (
	"fmt"
	"net/http"
	"net/url"

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

	user, err := h.domain.SaveToken(ctx, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	redirectURL := fmt.Sprintf(
		"http://local.fidelity.com:%s/client?id=%s&name=%s&email=%s",
		h.ClientPort,
		url.QueryEscape(user.ID),
		url.QueryEscape(user.Name),
		url.QueryEscape(user.Email),
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
