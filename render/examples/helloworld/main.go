package main
import (
	"log"
	"net/http"
	"hslam.com/git/x/mux"
	"hslam.com/git/x/handler/render"
)
func main() {
	r:=render.NewRender()
	m := mux.New()
	m.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.Text(w,req,"Hello world",http.StatusOK)
	}).All()
	log.Fatal(http.ListenAndServe(":8080", m))
}