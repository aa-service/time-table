package event

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aa-service/time-table/database"
	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetNoData(t *testing.T) {
	assert := assert.New(t)
	db := database.New("file:get?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)
	r := gin.New()
	w := httptest.NewRecorder()

	r.GET("/", get(opts))
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusOK)
	assert.JSONEq(`{"status":"ok", "data": []}`, w.Body.String())
}

func TestGetNoDataUUID(t *testing.T) {
	assert := assert.New(t)
	db := database.New("file:get?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)
	r := gin.New()
	w := httptest.NewRecorder()
	id := uuid.NewV4().String()

	r.Use(func(c *gin.Context) {
		c.Set("uuid", id)
	})
	r.GET("/:uuid", get(opts))

	req, _ := http.NewRequest("GET", "/"+id, nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusNotFound)
	assert.JSONEq(`{"status":"ko", "error": "not found"}`, w.Body.String())
}

func TestGetOK(t *testing.T) {
	assert := assert.New(t)
	db := database.New("file:get?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)
	r := gin.New()
	w := httptest.NewRecorder()
	id1 := uuid.NewV4().String()
	id2 := uuid.NewV4().String()

	db.Create(&models.Event{
		UUID: id1,
	})
	db.Create(&models.Event{
		UUID: id2,
	})

	r.Use(func(c *gin.Context) {
		c.Set("uuid", id1)
	})
	r.GET("/:uuid", get(opts))

	req, _ := http.NewRequest("GET", "/"+id1, nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusOK)
	// TODO: test json result

	// return all
	req, _ = http.NewRequest("GET", "/"+id1, nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusOK)
	// TODO: test json result
}
