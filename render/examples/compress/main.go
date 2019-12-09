package main
import (
	"log"
	"net/http"
	"github.com/hslam/mux"
	"github.com/hslam/handler/render"
)
func main() {
	r:=render.NewRender()
	r.GzipAll().DeflateAll().Charset("utf-8")
	m := mux.New()
	m.HandleFunc("/compress", func(w http.ResponseWriter, req *http.Request) {
		r.Body(w,req,[]byte("compress"),http.StatusOK)
	}).GET().POST().HEAD()
	log.Fatal(http.ListenAndServe(":8080", m))
}