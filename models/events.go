package models

import (
	"time"

	"gorm.io/gorm"
)

// Event contains event data
type Event struct {
	gorm.Model
	UUID string `gorm:"index,unique"`
	User User
	Day  byte
	Time time.Time
	Name string
}
