package event

import (
	"net/http"

	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

func get(opts *options.Options) gin.HandlerFunc {
	return func(c *gin.Context) {

		if uuid := c.GetString("uuid"); uuid == "" {
			// return uuid only
			var event models.User
			// todo: get events only for current user
			result := opts.DB().First(&event, "UUID = ?", uuid)
			//
			if result.Error != nil || result.RowsAffected == 0 {
				c.JSON(
					http.StatusNotFound,
					gin.H{
						"status": "ko",
						"error":  "not found",
					},
				)
				return
			}
			//
			c.JSON(
				http.StatusOK,
				gin.H{
					"status": "ok",
					"data":   event,
				},
			)
		}

		//
		// result := opts.DB()
		c.JSON(http.StatusOK, gin.H{"status": "getting events"})
	}
}
