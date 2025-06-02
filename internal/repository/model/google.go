package model

import (
	"time"

	"gorm.io/gorm"
)

type GoogleToken struct {
	gorm.Model
	UserEmail    string    `gorm:"unique;not null;index"`
	AccessToken  string    `gorm:"not null;index"`
	RefreshToken string    `gorm:"not null;index"`
	Expiry       time.Time `gorm:"not null;index"`
}
