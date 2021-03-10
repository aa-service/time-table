package database

import (
	"github.com/aa-service/time-table/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	ModeNormal = 0
	ModeDebug  = 1
)

func New(url string, mode int) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// prepare the db
	_ = db.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.UserToken{},
	)

	if mode == ModeDebug {
		return db.Debug()
	}

	return db
}
