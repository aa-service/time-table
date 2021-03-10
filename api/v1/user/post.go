package user

import (
	"net/http"

	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func post(opts *options.Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ko",
				"error":  "bad request",
			})
			return
		}

		user.UUID = uuid.NewV4().String()

		result := opts.DB().Create(&user)
		if result.Error != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status": "ko",
					"error":  "bad request",
				},
			)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   user.UUID,
		})
	}
}
