package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aa-service/time-table/database"
	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetKO1(t *testing.T) {
	assert := assert.New(t)
	r := gin.New()
	w := httptest.NewRecorder()
	db := database.New("file:get?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)

	r.GET("/", get(opts))

	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusNotFound)
	assert.JSONEq(`{"status":"ko","error":"not found"}`, w.Body.String())
}

func TestGetKO2(t *testing.T) {
	assert := assert.New(t)
	r := gin.New()
	w := httptest.NewRecorder()
	db := database.New("file:get?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)

	r.GET("/", func(c *gin.Context) {
		c.Set("uuid", "foobar")
	}, get(opts))

	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusNotFound)
	assert.JSONEq(`{"status":"ko","error":"not found"}`, w.Body.String())
}

func TestGetOK(t *testing.T) {
	assert := assert.New(t)
	r := gin.New()
	w := httptest.NewRecorder()
	db := database.New("file:get?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)

	user := models.User{
		UUID: "foobar",
	}
	_ = db.Create(&user)

	r.GET("/", func(c *gin.Context) {
		c.Set("uuid", "foobar")
	}, get(opts))

	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusOK)
	// assert.JSONEq(`{"status":"ok"}`, w.Body.String())
}
