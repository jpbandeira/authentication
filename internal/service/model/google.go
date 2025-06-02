package model

import "time"

type GoogleToken struct {
	ID           string
	UserEmail    string
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
