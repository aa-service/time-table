package database_test

import (
	"testing"

	"github.com/aa-service/time-table/database"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	db := database.New("file:db?mode=memory&cache=shared", database.ModeDefault)
	assert.NotNil(db)

	db = database.New("file:db?mode=memory&cache=shared", database.ModeDebug)
	assert.NotNil(db)
}
