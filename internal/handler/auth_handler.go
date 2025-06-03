package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jp/authentication/internal/pkg/dto"
)

func (h *AuthHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var credsDTO dto.Creds
	if err := c.ShouldBindJSON(&credsDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	token, err := h.domain.Login(ctx, credsDTO.Email, credsDTO.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
