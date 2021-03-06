package main

import (
	"github.com/gin-gonic/gin"
)

func initAPI(router *gin.RouterGroup, auth gin.HandlerFunc) {
	NewV1(router.Group("v1"), auth)
}
