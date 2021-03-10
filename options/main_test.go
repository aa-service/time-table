package options_test

import (
	"testing"

	"github.com/aa-service/time-table/options"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestErr1(t *testing.T) {
	assert := assert.New(t)

	out, err := options.New()

	assert.Nil(out)
	assert.NotNil(err)
}

func TestErr2(t *testing.T) {
	assert := assert.New(t)

	out, err := options.New("foo", "bar", "baz")

	assert.Nil(out)
	assert.NotNil(err)
}

func TestOk(t *testing.T) {
	assert := assert.New(t)

	ptr := &gorm.DB{}
	out, err := options.New("foo", "bar", ptr, "baz")

	assert.Nil(err)
	assert.NotNil(out)
	assert.Equal(ptr, out.DB())
}
