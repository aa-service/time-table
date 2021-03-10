package event

import (
	"net/http"

	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

func get(opts *options.Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		if uuid := c.GetString("uuid"); uuid != "" {
			// return uuid only
			var event models.Event
			// todo: get events only for current user
			result := opts.DB().First(&event, "UUID = ?", uuid)
			//
			if result.RowsAffected == 0 {
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
			return
		}

		//
		var events []models.Event
		_ = opts.DB().Find(&events)
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": "ok",
				"data":   events,
			},
		)
	}
}
