package domain

import (
	"context"

	"github.com/jp/authentication/internal/pkg/auth_hash"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func (d *domain) Register(ctx context.Context, user User) error {
	hashed, err := auth_hash.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed

	return d.db.Register(ctx, user)
}
