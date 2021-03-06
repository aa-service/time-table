package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewV1(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
}
