package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aa-service/time-table/api"
	"github.com/aa-service/time-table/database"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB
var router *gin.Engine

const defaultPort = "8080"

func init() {
	db = database.New("database.db", database.ModeDebug)

	router = gin.Default()
	// initAPI(router.Group("api"), authorizator)
	options, err := options.New(db)
	if err != nil {
		panic(err)
	}

	_ = api.Mount(router.Group("api"), options)
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
