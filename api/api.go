package api

import (
	v1 "github.com/aa-service/time-table/api/v1"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

func New(router *gin.RouterGroup, opts *options.Options) {
	v1.New(router.Group("v1"), opts)
}
