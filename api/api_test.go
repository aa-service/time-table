package api_test

import (
	"testing"

	"github.com/aa-service/time-table/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMount(t *testing.T) {
	r := gin.New()
	g := r.Group("/")
	mounted := api.Mount(g, nil)
	//
	assert.NotNil(t, mounted)
	assert.NotNil(t, mounted["v1"])
}
