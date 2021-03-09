package api

import (
	v1 "github.com/aa-service/time-table/api/v1"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

func Mount(router *gin.RouterGroup, opts *options.Options) (groups map[string]*gin.RouterGroup) {
	groups = map[string]*gin.RouterGroup{
		"v1": router.Group("v1"),
	}
	v1.Mount(groups["v1"], opts)
	//
	return
}
