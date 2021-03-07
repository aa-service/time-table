package event

import (
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

func Mount(router *gin.RouterGroup, opts *options.Options) {
	router.GET("/*uuid", opts.Authorizator(), get(opts))
	router.POST("/", opts.Authorizator(), post(opts))
	// router.DELETE("/:uuid", opts.Authorizator(), delete(opts))
}
