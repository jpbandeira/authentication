package handler

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	dtoModel "github.com/jp/authentication/internal/model"
	rModel "github.com/jp/authentication/internal/repository/model"

	"github.com/jp/authentication/internal/repository"
	"github.com/jp/authentication/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	repo *repository.UserRepository
	auth *service.AuthService
}

func NewAuthHandler(r *repository.UserRepository, a *service.AuthService) *AuthHandler {
	return &AuthHandler{repo: r, auth: a}
}

func (h *AuthHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	var user dtoModel.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	hashed, err := h.auth.HashPassword(user.Password)
	if err != nil {
		log.Printf("[%v] erro hash: %v", ctx.Value("req_id"), err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno"})
		return
	}
	user.Password = hashed

	userRepo := &rModel.User{
		UUID:     uuid.NewString(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := h.repo.Create(ctx, userRepo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao cadastrar"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": userRepo.UUID})
}

func (h *AuthHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	user, err := h.repo.FindByEmail(ctx, creds.Email)
	if err != nil || !h.auth.CheckPassword(user.Password, creds.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciais inválidas"})
		return
	}
	token, err := h.auth.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao gerar token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
