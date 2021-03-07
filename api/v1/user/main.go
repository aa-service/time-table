package user

import (
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

func Mount(router *gin.RouterGroup, opts *options.Options) {
	router.POST("/", post(opts))
	router.GET("/:uuid", opts.Authorizator(), get(opts))
	router.DELETE("/:uuid", opts.Authorizator(), delete(opts))
}
