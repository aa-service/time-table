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
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestLoginKO(t *testing.T) {
	assert := assert.New(t)
	r := gin.New()
	w := httptest.NewRecorder()
	opts := &options.Options{}

	r.POST("/", login(opts))

	req, _ := http.NewRequest("POST", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusBadRequest)
	assert.JSONEq(`{"status":"ko","error":"bad request"}`, w.Body.String())
}

func TestLoginNoUser(t *testing.T) {
	assert := assert.New(t)
	r := gin.New()
	w := httptest.NewRecorder()
	db := database.New("file:post?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)

	r.POST("/", login(opts))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"email": "foobar"}`)))
	r.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusNotFound)
	assert.JSONEq(`{"status":"ko","error":"not found"}`, w.Body.String())
}

func TestLoginOK(t *testing.T) {
	assert := assert.New(t)
	r := gin.New()
	w := httptest.NewRecorder()
	db := database.New("file:post?mode=memory&cache=shared", database.ModeDebug)
	opts, _ := options.New(db)
	user := models.User{
		UUID:  uuid.NewV4().String(),
		Email: "foobazbaz",
	}
	_ = db.Create(&user)

	r.POST("/", login(opts))

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(
		[]byte(`{"email": "foobarbaz"}`)),
	)
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
	uToken := models.UserToken{}
	_ = db.First(&uToken, "token = ?", js.Data)

	assert.Equal(user.ID, uToken.UserID)
}
