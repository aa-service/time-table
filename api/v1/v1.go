package v1

import (
	"github.com/aa-service/time-table/api/v1/event"
	"github.com/aa-service/time-table/api/v1/middleware"
	"github.com/aa-service/time-table/api/v1/user"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

func Mount(router *gin.RouterGroup, options *options.Options) (groups map[string]*gin.RouterGroup) {
	groups = map[string]*gin.RouterGroup{
		"user":  router.Group("/user"),
		"event": router.Group("/event"),
	}
	//
	router.Use(middleware.ChkUUID())
	auth := middleware.Auth(options)
	//
	user.Mount(groups["user"], options, auth)
	event.Mount(groups["event"], options, auth)
	//
	return
}
