package repository

import (
	"context"

	rmodel "github.com/jp/authentication/internal/repository/model"
	smodel "github.com/jp/authentication/internal/service/model"
)

func (s *Repository) SaveGoogleToken(ctx context.Context, googleToken smodel.GoogleToken) (string, error) {
	dbModel := rmodel.GoogleToken{
		UserEmail:    googleToken.UserEmail,
		AccessToken:  googleToken.AccessToken,
		RefreshToken: googleToken.RefreshToken,
		Expiry:       googleToken.Expiry,
	}

	if err := s.db.WithContext(ctx).Where(rmodel.GoogleToken{UserEmail: googleToken.UserEmail}).Assign(dbModel).FirstOrCreate(&dbModel).Error; err != nil {
		return "", err
	}

	return dbModel.UserEmail, nil
}
