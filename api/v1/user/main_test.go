package user_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aa-service/time-table/api/v1/user"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Test that paths are mounted and auth is required when needed
func TestMount(t *testing.T) {
	for _, test := range []struct {
		in     string
		method string
		auth   bool
		code   int
	}{
		{"/", "POST", false, http.StatusBadRequest},
		{"/login", "POST", false, http.StatusBadRequest},
		{"/2928", "GET", true, http.StatusNotFound},
		{"/2928", "DELETE", true, http.StatusNotFound},
	} {
		opts, _ := options.New(&gorm.DB{})
		r := gin.New()
		user.Mount(r.Group(""), opts, func(c *gin.Context) {
			assert.True(t, test.auth)
		})
		//
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(test.method, test.in, nil)
		r.ServeHTTP(w, req)
		//
		assert.Equal(t, test.code, w.Code)
	}
}
