package database

import (
	"github.com/aa-service/time-table/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(url string, debug bool) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// prepare the db
	db.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.UserToken{},
	)

	if debug {
		return db.Debug()
	}

	return db
}
