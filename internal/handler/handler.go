package handler

import "github.com/jp/authentication/internal/domain"

type AuthHandler struct {
	domain     domain.Domain
	ClientPort string
}

func NewAuthHandler(d domain.Domain, clietPort string) *AuthHandler {
	return &AuthHandler{domain: d, ClientPort: clietPort}
}
