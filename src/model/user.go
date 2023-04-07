package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:64;not null"`
	Password string `gorm:"size:128;not null"`
}
