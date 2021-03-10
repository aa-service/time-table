package event_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aa-service/time-table/api/v1/event"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestMountAndAuthCheck(t *testing.T) {
	assert := assert.New(t)

	for _, test := range []struct {
		path   string
		method string
		auth   bool
	}{
		{"/", "GET", true},
		{"/", "POST", true},
		{"/" + uuid.NewV4().String(), "GET", true},
		{"/" + uuid.NewV4().String(), "DELETE", true},
	} {
		called := false
		r := gin.New()
		group := r.Group("/")
		event.Mount(group, nil, func(c *gin.Context) {
			called = true
			c.Abort()
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(test.method, test.path, nil)
		r.ServeHTTP(w, req)
		//
		assert.Equal(test.auth, called)
	}
}
