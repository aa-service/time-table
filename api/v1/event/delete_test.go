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

func TestDeleteNoData(t *testing.T) {
	assert := assert.New(t)
	db := database.New("file:delete?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)
	r := gin.New()
	w := httptest.NewRecorder()
	id := uuid.NewV4().String()
	event := models.Event{UUID: id}

	_ = db.Create(&event)

	r.DELETE("/", delete(opts))

	req, _ := http.NewRequest("DELETE", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusNotFound)
	assert.JSONEq(`{"status":"ko", "error": "not found"}`, w.Body.String())
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	db := database.New("file:delete?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)
	r := gin.New()
	w := httptest.NewRecorder()
	id := uuid.NewV4().String()
	event := models.Event{UUID: id}

	_ = db.Create(&event)

	r.DELETE("/", func(c *gin.Context) {
		c.Set("uuid", id)
	}, delete(opts))

	req, _ := http.NewRequest("DELETE", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusOK)
	assert.JSONEq(`{"status":"ok"}`, w.Body.String())

	event = models.Event{}
	_ = db.First(&event, "uuid = ?", id)

	assert.Empty(event.UUID)
}
