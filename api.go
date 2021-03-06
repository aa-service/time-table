package main

import (
	"github.com/gin-gonic/gin"
)

func initAPI(router *gin.RouterGroup) {
	NewV1(router.Group("v1"))
}
