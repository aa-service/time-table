package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aa-service/time-table/api/v1/middleware"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var opts *options.Options

const dbURL = "file::memory:?cache=shared"

func init() {
	db, _ = gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	opts, _ = options.New(db.Debug())
}

func TestAuthorizator(t *testing.T) {
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
			c.String(http.StatusOK, c.GetString("tested"))
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
