package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aa-service/time-table/api/v1/middleware"
	"github.com/aa-service/time-table/database"
	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB
var opts *options.Options

const dbURL = "file:foobar?mode=memory&cache=shared"

func init() {
	db = database.New(dbURL, database.ModeDebug)
	opts, _ = options.New(db)
}

func TestAuthorizatorKO(t *testing.T) {
	for _, test := range []struct {
		headerName string
		in         string
		outCode    int
		nextEd     bool
	}{
		{"foo", "", 401, false},
		{"foo", "foo", 401, false},
		{"Authorization", "foo", 401, false},
	} {
		nextEd := false
		r := gin.New()
		r.Use(middleware.Auth(opts))
		r.GET("/", func(c *gin.Context) {
			nextEd = true
			c.String(http.StatusOK, "ok")
		})
		//
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set(test.headerName, test.in)
		r.ServeHTTP(w, req)
		//
		assert.Equal(t, test.outCode, w.Code)
		assert.Equal(t, test.nextEd, nextEd)
	}
}

// A valid token but no user, must fail
func TestAuthorizatorKONoUser(t *testing.T) {
	called := false

	// create auth token
	uToken := models.UserToken{
		UserID: 12,
		Token:  uuid.NewV4().String(),
	}
	_ = opts.DB().Create(&uToken)

	// router
	r := gin.New()
	r.Use(middleware.Auth(opts))
	r.GET("/", func(c *gin.Context) {
		called = true
		c.String(http.StatusOK, "ok")
	})

	// finally
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", uToken.Token)
	r.ServeHTTP(w, req)

	//
	assert.False(t, called)
}

func TestAuthorizatorOK(t *testing.T) {
	called := false
	// create user
	user := models.User{
		UUID: uuid.NewV4().String(),
	}
	_ = opts.DB().Create(&user)

	// create auth token
	uToken := models.UserToken{
		UserID: user.ID,
		Token:  uuid.NewV4().String(),
	}
	_ = opts.DB().Create(&uToken)

	// router
	r := gin.New()
	r.Use(middleware.Auth(opts))
	r.GET("/", func(c *gin.Context) {
		called = true
		c.String(http.StatusOK, "ok")
		// check response
		raw, exists := c.Get("user")
		assert.True(t, exists)
		u := raw.(models.User)
		assert.Equal(t, user.ID, u.ID)
		assert.Equal(t, user.UUID, u.UUID)
	})

	// finally
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", uToken.Token)
	r.ServeHTTP(w, req)
	//
	assert.True(t, called)
}
