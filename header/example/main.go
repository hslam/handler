package main

import (
	"github.com/hslam/handler/header"
	"github.com/hslam/mux"
	"log"
	"net/http"
)

func main() {
	m := mux.New()
	m.Use(func(w http.ResponseWriter, r *http.Request) {
		header.SetHeader(w, header.AccessControlAllowOrigin, "*")
	})
	m.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).All()
	log.Fatal(http.ListenAndServe(":8080", m))
}
