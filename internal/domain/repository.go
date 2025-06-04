package domain

import (
	"context"
)

type Repository interface {
	Register(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (User, error)

	SaveGoogleToken(ctx context.Context, googleToken GoogleToken) error
	GetGoogleToken(ctx context.Context, email string) error
}
