package main

import (
	"io"
	"net/http"

	"gorm.io/gorm"
)

// User contains the user data
type User struct {
	gorm.Model
	UUID       string `gorm:"index"`
	Name       string
	TZ         string
	UserTokens []UserToken
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "User")
}
