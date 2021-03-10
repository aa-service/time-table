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

func TestDeleteKO1(t *testing.T) {
	assert := assert.New(t)
	r := gin.New()
	w := httptest.NewRecorder()
	db := database.New("file:post?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)

	r.DELETE("/", delete(opts))

	req, _ := http.NewRequest("DELETE", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusNotFound)
	assert.JSONEq(`{"status":"ko","error":"not found"}`, w.Body.String())
}

func TestDeleteKO2(t *testing.T) {
	assert := assert.New(t)
	r := gin.New()
	w := httptest.NewRecorder()
	db := database.New("file:post?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)

	r.DELETE("/", func(c *gin.Context) {
		c.Set("uuid", "foobar")
	}, delete(opts))

	req, _ := http.NewRequest("DELETE", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusNotFound)
	assert.JSONEq(`{"status":"ko","error":"not found"}`, w.Body.String())
}

func TestDeleteOK(t *testing.T) {
	assert := assert.New(t)
	r := gin.New()
	w := httptest.NewRecorder()
	db := database.New("file:post?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)

	user := models.User{
		UUID: "foobar",
	}
	_ = db.Create(&user)

	r.DELETE("/", func(c *gin.Context) {
		c.Set("uuid", "foobar")
	}, delete(opts))

	req, _ := http.NewRequest("DELETE", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusOK)
	assert.JSONEq(`{"status":"ok"}`, w.Body.String())

	var u2 models.User
	_ = db.First(&u2, "uuid = ?", user.UUID)

	assert.Empty(u2.UUID)
}
