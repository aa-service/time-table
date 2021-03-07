package user

import (
	"net/http"

	"github.com/aa-service/time-table/options"
	"github.com/brugnara/uuidchecker"
	"github.com/gin-gonic/gin"
)

func Mount(router *gin.RouterGroup, opts *options.Options) {
	router.POST("/", post(opts))
	router.GET("/:uuid", chkUUID(), get(opts))
	router.DELETE("/:uuid", chkUUID(), delete(opts))
}

func chkUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("uuid", c.Param("uuid"))
		if !uuidchecker.IsValid(c.Keys["uuid"].(string)) {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status": "ko",
					"error":  "invalid UUID provided",
				},
			)
			c.Abort()
			return
		}
		c.Next()
	}
}
