package repository

import (
	"context"
	"log"

	"github.com/jp/authentication/internal/domain"
	rmodel "github.com/jp/authentication/internal/repository/model"
)

func (db *GormRepository) SaveGoogleToken(ctx context.Context, googleToken domain.GoogleToken) error {
	dbModel := rmodel.GoogleToken{
		UserEmail:    googleToken.UserEmail,
		AccessToken:  googleToken.AccessToken,
		RefreshToken: googleToken.RefreshToken,
		Expiry:       googleToken.Expiry,
	}

	if err := db.WithContext(ctx).Where(rmodel.GoogleToken{UserEmail: googleToken.UserEmail}).Assign(dbModel).FirstOrCreate(&dbModel).Error; err != nil {
		return err
	}

	return nil
}

func (db *GormRepository) GetGoogleToken(ctx context.Context, email string) error {
	var googleToken rmodel.GoogleToken
	if err := db.WithContext(ctx).Where("user_email = ?", email).First(&googleToken).Error; err != nil {
		log.Printf("[%v] erro ao buscar usuário por email: %v", ctx.Value("req_id"), err)
		return err
	}

	return nil
}
