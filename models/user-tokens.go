package models

import "gorm.io/gorm"

// UserToken contains tokens usable by an user
type UserToken struct {
	gorm.Model
	UserID uint
	Token  string `gorm:"uniqueIndex"`
}
