package models

import (
	"gorm.io/gorm"
)

// User contains the user data
type User struct {
	gorm.Model
	Email string `gorm:"uniqueIndex" json:"email"`
	UUID  string `gorm:"uniqueIndex" json:"uuid"`
	Name  string `json:"name"`
	TZ    string `json:"tz"`
	// UserTokens []UserToken
}
