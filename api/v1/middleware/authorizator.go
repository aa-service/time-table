package middleware

import (
	"log"
	"net/http"

	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
)

type header struct {
	Token string `header:"Authorization"`
}

func Auth(opts *options.Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("requested auth")

		var hdr header
		if err := c.ShouldBindHeader(&hdr); err != nil || hdr.Token == "" {
			log.Println("no auth header :/")
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		var user models.User
		var uToken models.UserToken

		result := opts.DB().
			Where("token = ?", hdr.Token).
			Find(&uToken)

		if result.Error != nil || result.RowsAffected == 0 {
			log.Println("no auth header :/")
			c.JSON(http.StatusUnauthorized, result.Error.Error())
			c.Abort()
			return
		}

		// find user
		result = opts.DB().
			Where("id = ?", uToken.UserID).
			Find(&user)

		if result.Error != nil || result.RowsAffected == 0 {
			log.Println("no auth header :/")
			c.JSON(http.StatusUnauthorized, result.Error.Error())
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
