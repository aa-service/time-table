package event

import (
	"net/http"

	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func post(opts *options.Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		var event models.Event

		if err := c.ShouldBindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ko",
				"error":  "bad request",
			})
			return
		}

		event.UUID = uuid.NewV4().String()

		result := opts.DB().Create(&event)
		if result.Error != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status": "ko",
					"error":  "uuid duplicated",
				},
			)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   event.UUID,
		})
	}
}
