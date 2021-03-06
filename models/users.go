package models

import (
	"gorm.io/gorm"
)

// User contains the user data
type User struct {
	gorm.Model
	Email  string      `gorm:"uniqueIndex" json:"email"`
	UUID   string      `gorm:"uniqueIndex" json:"uuid"`
	Name   string      `json:"name"`
	TZ     string      `json:"tz"`
	Tokens []UserToken `gorm:"constraint:OnDelete:CASCADE"`
	Events []Event     `gorm:"constraint:OnDelete:CASCADE"`
}
