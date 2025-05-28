package service

import (
	"time"

	"github.com/jp/authentication/internal/config"
	"github.com/jp/authentication/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	secret string
}

func NewAuthService() *AuthService {
	return &AuthService{secret: config.JWTSecret()}
}

func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func (s *AuthService) CheckPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (s *AuthService) GenerateToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"name":  user.Name,
		"email": user.Email,
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}
