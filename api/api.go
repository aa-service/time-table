package api

import (
	v1 "github.com/aa-service/time-table/api/v1"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

type Api struct {
	router  *gin.RouterGroup
	options *options.Options
}

func New(router *gin.RouterGroup, opts *options.Options) (*Api, error) {
	// NewV1(router.Group("v1"), auth)
	api := &Api{router, opts}
	//
	v1.New(router.Group("v1"), opts)
	//
	return api, nil
}
