package domain

import (
	"context"
)

type Domain interface {
	Register(context.Context, User) error
	FindByEmail(context.Context, string) (User, error)

	Login(context.Context, string, string) (string, error)

	SaveToken(ctx context.Context, code string) (User, error)
}

type domain struct {
	db     Repository
	secret string
}

func NewDomain(db Repository, jwtSecret string) Domain {
	return &domain{
		db:     db,
		secret: jwtSecret,
	}
}
