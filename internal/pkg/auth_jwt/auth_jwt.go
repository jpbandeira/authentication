package auth_jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(id, name, email string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   id,
		"name":  name,
		"email": email,
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
