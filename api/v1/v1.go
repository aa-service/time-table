package v1

import (
	"github.com/aa-service/time-table/api/v1/event"
	"github.com/aa-service/time-table/api/v1/middleware"
	"github.com/aa-service/time-table/api/v1/user"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

type V1 struct{}

func New(router *gin.RouterGroup, options *options.Options) {
	router.Use(middleware.ChkUUID())
	auth := middleware.Auth(options)
	//
	user.Mount(router.Group("/user"), options, auth)
	event.Mount(router.Group("/event"), options, auth)
}
