package main

import (
	"github.com/hslam/handler/render"
	"github.com/hslam/mux"
	"log"
	"net/http"
)

func main() {
	r := render.NewRender()
	r.GzipAll().DeflateAll().Charset("utf-8")
	m := mux.New()
	m.HandleFunc("/compress", func(w http.ResponseWriter, req *http.Request) {
		r.Body(w, req, []byte("compress"), http.StatusOK)
	}).GET().POST().HEAD()
	log.Fatal(http.ListenAndServe(":8080", m))
}
