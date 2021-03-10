package user

import (
	"net/http"

	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func login(opts *options.Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var uToken models.UserToken

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "ko",
				"error":  "bad request",
			})
			return
		}

		result := opts.DB().First(&user)

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

		// create new token
		uToken.Token = uuid.NewV4().String()
		uToken.UserID = user.ID

		// set token to user
		opts.DB().Create(&uToken)

		// print token
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   uToken.Token,
		})
	}
}
