package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc(fmt.Sprintf("/%s/", APIVs), hIndex)
	//
	http.Handle(fmt.Sprintf("/%s/event/", APIVs), &Event{})
	http.Handle(fmt.Sprintf("/%s/user/", APIVs), &User{})
}

// APIListen starts the api
func APIListen(port string) {
	if port == "" {
		port = "8080"
	}
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}

// handlers
func hIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index")
}
