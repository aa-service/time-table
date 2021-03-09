package v1_test

import (
	"testing"

	v1 "github.com/aa-service/time-table/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMount(t *testing.T) {
	r := gin.New()
	mounted := v1.Mount(r.Group("/"), nil)
	assert.NotNil(t, mounted)
	assert.NotNil(t, mounted["user"])
	assert.NotNil(t, mounted["event"])
}
