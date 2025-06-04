package repository

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jp/authentication/internal/domain"
	"github.com/jp/authentication/internal/repository/model"
	rmodel "github.com/jp/authentication/internal/repository/model"
)

func (db *GormRepository) Register(ctx context.Context, user domain.User) error {
	userRepo := &model.User{
		UUID:     uuid.NewString(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := db.WithContext(ctx).Create(userRepo).Error; err != nil {
		log.Printf("[%v] erro ao criar usuário: %v", ctx.Value("req_id"), err)
		return err
	}
	return nil
}

func (db *GormRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var user rmodel.User
	if err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		log.Printf("[%v] erro ao buscar usuário por email: %v", ctx.Value("req_id"), err)
		return domain.User{}, err
	}

	return domain.User{
		ID:       user.UUID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
