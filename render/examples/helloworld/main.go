package main

import (
	"github.com/hslam/handler/render"
	"github.com/hslam/mux"
	"log"
	"net/http"
)

func main() {
	r := render.NewRender()
	m := mux.New()
	m.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.Text(w, req, "Hello world", http.StatusOK)
	}).All()
	log.Fatal(http.ListenAndServe(":8080", m))
}
