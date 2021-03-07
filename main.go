package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aa-service/time-table/api"
	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var router *gin.Engine

const defaultPort = "8080"

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router = gin.Default()
	// initAPI(router.Group("api"), authorizator)
	options, err := options.New(db.Debug())
	if err != nil {
		panic(err)
	}
	api.New(router.Group("api"), options)

	// prepare the db
	db.AutoMigrate(
		&models.User{},
		&models.Event{},
		&models.UserToken{},
	)
}

func main() {
	log.Fatal(router.Run(getPort()))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return fmt.Sprintf(":%s", port)
}
