package event

import (
	"log"
	"net/http"

	"github.com/aa-service/time-table/options"
	"github.com/brugnara/uuidchecker"
	"github.com/gin-gonic/gin"
)

func Mount(router *gin.RouterGroup, opts *options.Options) {
	router.GET("/*uuid", opts.Authorizator(), func(c *gin.Context) {
		uuid := c.Param("uuid")
		if uuidchecker.IsValid(uuid) {
			log.Println("Valid uuid!")
		}
		c.JSON(http.StatusOK, gin.H{"status": "getting event: " + uuid})
	})
}
