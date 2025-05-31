package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     string `gorm:"unique;not null;index"`
	Email    string `gorm:"uniqueIndex;not null"`
	Name     string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
