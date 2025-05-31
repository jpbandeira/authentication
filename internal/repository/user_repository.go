package repository

import (
	"context"
	"log"

	"github.com/jp/authentication/internal/repository/model"
	rmodel "github.com/jp/authentication/internal/repository/model"
	smodel "github.com/jp/authentication/internal/service/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		log.Printf("[%v] erro ao criar usuário: %v", ctx.Value("req_id"), err)
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*smodel.User, error) {
	var user rmodel.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		log.Printf("[%v] erro ao buscar usuário por email: %v", ctx.Value("req_id"), err)
		return nil, err
	}
	return &smodel.User{
		ID:       user.UUID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
