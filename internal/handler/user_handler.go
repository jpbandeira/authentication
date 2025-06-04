package handler

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/jp/authentication/internal/domain"
	"github.com/jp/authentication/internal/pkg/dto"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	var user dto.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	userDomain := domain.User{
		ID:       uuid.NewString(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	if err := h.domain.Register(ctx, userDomain); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": userDomain})
}

func (h *AuthHandler) FindByEmail(c *gin.Context) {
	ctx := c.Request.Context()
	email := c.Param("email")
	if len(strings.TrimSpace(email)) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid email"})
		return
	}

	user, err := h.domain.FindByEmail(ctx, email)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find user", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}
