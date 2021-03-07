package v1

import (
	"github.com/aa-service/time-table/api/v1/event"
	"github.com/aa-service/time-table/api/v1/user"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

type V1 struct{}

func New(router *gin.RouterGroup, options *options.Options) {
	user.Mount(router.Group("/user"), options)
	event.Mount(router.Group("/event"), options)
}
