package models

import (
	"time"

	"gorm.io/gorm"
)

// Event contains event data
type Event struct {
	gorm.Model
	UserID uint
	UUID   string    `gorm:"uniqueIndex" json:"uuid"`
	Day    byte      `json:"day"`
	Time   time.Time `json:"time"`
	Name   string    `json:"name"`
}
