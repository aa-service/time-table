package user

import (
	"net/http"

	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

func get(opts *options.Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		result := opts.DB().First(&user, "UUID = ?", c.MustGet("uuid"))

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

		c.JSON(
			http.StatusOK,
			gin.H{
				"status": "ok",
				"data":   user,
			},
		)
	}
}
