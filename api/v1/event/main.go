package event

import (
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

func Mount(
	router *gin.RouterGroup,
	opts *options.Options,
	auth gin.HandlerFunc,
) {
	router.GET("/:uuid", auth, get(opts))
	router.GET("/", auth, get(opts))
	router.POST("/", auth, post(opts))
	router.DELETE("/:uuid", auth, delete(opts))
}
