package v1

import (
	"github.com/aa-service/time-table/api/v1/event"
	"github.com/aa-service/time-table/api/v1/user"
	"github.com/aa-service/time-table/options"
	"github.com/brugnara/uuidchecker"
	"github.com/gin-gonic/gin"
)

type V1 struct{}

func New(router *gin.RouterGroup, options *options.Options) {
	router.Use(chkUUID())
	//
	user.Mount(router.Group("/user"), options)
	event.Mount(router.Group("/event"), options)
}

func chkUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if uuid := c.Param("uuid"); uuidchecker.IsValid(uuid) {
			c.Set("uuid", uuid)
		}
		c.Next()
	}
}
