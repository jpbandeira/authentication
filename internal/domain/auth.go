package domain

import (
	"context"

	"github.com/jp/authentication/internal/pkg/auth_hash"
	"github.com/jp/authentication/internal/pkg/auth_jwt"
)

func (d *domain) Login(ctx context.Context, email, password string) (string, error) {
	user, err := d.FindByEmail(ctx, email)
	if err != nil || !auth_hash.CheckPassword(user.Password, password) {
		return "", err
	}

	token, err := auth_jwt.GenerateToken(user.ID, user.Name, user.Email, d.secret)
	if err != nil {
		return "", err
	}

	return token, nil
}
