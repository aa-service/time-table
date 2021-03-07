package user

import (
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

func Mount(
	router *gin.RouterGroup,
	opts *options.Options,
	auth gin.HandlerFunc,
) {
	router.POST("/", post(opts))
	router.POST("/login", login(opts))
	router.GET("/:uuid", auth, get(opts))
	router.DELETE("/:uuid", auth, delete(opts))
}
