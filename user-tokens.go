package main

import "gorm.io/gorm"

// UserToken contains tokens usable by an user
type UserToken struct {
	gorm.Model
	User  User
	Token string `gorm:"index"`
}
