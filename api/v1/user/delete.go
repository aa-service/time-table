package user

import (
	"net/http"

	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

func delete(opts *options.Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := opts.DB().
			Unscoped().
			Delete(models.User{}, "uuid = ?", c.MustGet("uuid"))

		if result.Error != nil {
			c.JSON(
				http.StatusOK,
				gin.H{
					"status": "ko",
					"error":  result.Error.Error(),
				},
			)
			return
		}

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

		c.JSON(
			http.StatusOK,
			gin.H{
				"status": "ok",
			},
		)
	}
}
