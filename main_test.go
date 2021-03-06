package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPort(t *testing.T) {
	asrt := assert.New(t)

	asrt.Equal(getPort(), ":"+defaultPort)

	os.Setenv("PORT", "9999")
	asrt.Equal(getPort(), ":9999")
}
