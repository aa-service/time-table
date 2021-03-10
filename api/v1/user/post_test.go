package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aa-service/time-table/database"
	"github.com/aa-service/time-table/models"
	"github.com/aa-service/time-table/options"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostNoData(t *testing.T) {
	assert := assert.New(t)
	db := database.New("file:post?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)
	r := gin.New()
	w := httptest.NewRecorder()

	r.POST("/", post(opts))

	req, _ := http.NewRequest("POST", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusBadRequest)
	assert.JSONEq(`{"status":"ko", "error": "bad request"}`, w.Body.String())
}

func TestPostData(t *testing.T) {
	assert := assert.New(t)
	db := database.New("file:post?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)
	r := gin.New()
	w := httptest.NewRecorder()

	r.POST("/", post(opts))
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"name":"name"}`)))
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusOK)

	js := struct {
		Status string `json:"status"`
		Data   string `json:"data"`
	}{}
	_ = json.NewDecoder(w.Body).Decode(&js)

	assert.Equal("ok", js.Status)
	assert.NotEmpty(js.Data)

	// check db for presence
	var user models.User
	_ = db.First(&user, "uuid = ?", js.Data)

	assert.Equal(js.Data, user.UUID)
}
