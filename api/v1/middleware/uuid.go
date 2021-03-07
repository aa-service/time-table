package middleware

import (
	"github.com/brugnara/uuidchecker"
	"github.com/gin-gonic/gin"
)

func ChkUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if uuid := c.Param("uuid"); uuidchecker.IsValid(uuid) {
			c.Set("uuid", uuid)
		}
		c.Next()
	}
}
