package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aa-service/time-table/api/v1/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestChkUUID(t *testing.T) {
	for _, test := range []struct {
		in  string
		out string
	}{
		{"/foobar", ""},
		{"/ff5f1af1-6b1c-4ef0-be13-765145aaf1a3junk", ""},
		{"/junkff5f1af1-6b1c-4ef0-be13-765145aaf1a3", ""},
		{"/ff5f1af1-6b1c-4ef0-be13-765145aaf1a3", "ff5f1af1-6b1c-4ef0-be13-765145aaf1a3"},
	} {
		r := gin.New()
		r.Use(middleware.ChkUUID())
		r.GET("/:uuid", func(c *gin.Context) {
			c.String(http.StatusOK, c.GetString("uuid"))
		})
		//
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.in, nil)
		r.ServeHTTP(w, req)
		//
		assert.Equal(t, test.out, w.Body.String())
	}
}
