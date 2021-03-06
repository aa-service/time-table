package main

import (
	"log"
	"net/http"

	"github.com/brugnara/uuidchecker"
	"github.com/gin-gonic/gin"
)

func NewV1(router *gin.RouterGroup, auth gin.HandlerFunc) {
	router.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "creating user"})
	})

	router.GET("/user/:uuid", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "getting user"})
	})

	router.DELETE("/user/:uuid", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "deleting user"})
	})

	//

	router.GET("/event/*uuid", auth, func(c *gin.Context) {
		uuid := c.Param("uuid")
		if uuidchecker.IsValid(uuid) {
			log.Println("Valid uuid!")
		}
		c.JSON(http.StatusOK, gin.H{"status": "getting event: " + uuid})
	})
}
