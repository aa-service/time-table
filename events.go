package main

import (
	"io"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// Event contains event data
type Event struct {
	gorm.Model
	UUID string `gorm:"index"`
	User User
	Day  byte
	Time time.Time
	Name string
}

func (e *Event) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Event")
}
