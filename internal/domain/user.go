package domain

import (
	"context"
	"errors"

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

func (d *domain) FindByEmail(ctx context.Context, email string) (User, error) {
	if email == "" {
		return User{}, errors.New("email invalido")
	}

	user, err := d.db.FindByEmail(ctx, email)
	if err != nil {
		if user.ID == "" {
			return User{}, errors.New("not found")
		}

		return User{}, err
	}

	return user, nil
}
