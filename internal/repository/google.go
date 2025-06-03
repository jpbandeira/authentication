package repository

import (
	"context"

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
