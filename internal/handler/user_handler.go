package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jp/authentication/internal/domain"
	"github.com/jp/authentication/internal/pkg/dto"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	var user dto.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	userDomain := domain.User{
		ID:       uuid.NewString(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	if err := h.domain.Register(ctx, userDomain); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao cadastrar"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": userDomain})
}
