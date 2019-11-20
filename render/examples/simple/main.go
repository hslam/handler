package main
import (
	"log"
	"net/http"
	"hslam.com/git/x/mux"
	"hslam.com/git/x/handler/render"
)
type Student struct {
	Name string
	Age int32
	Address string
}
func main() {
	r:=render.NewRender()
	r.GzipAll().DeflateAll().Charset("utf-8")
	router := mux.New()
	router.HandleFunc("/text", func(w http.ResponseWriter, req *http.Request) {
		r.Text(w,req,"Hello world",http.StatusOK)
	}).All()
	router.HandleFunc("/raw", func(w http.ResponseWriter, req *http.Request) {
		r.Body(w,req,[]byte("raw data"),http.StatusOK)
	}).All()
	router.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w,req,Student{"Mort Huang",18,"Earth"},http.StatusOK)
	}).All()
	router.HandleFunc("/xml", func(w http.ResponseWriter, req *http.Request) {
		r.XML(w,req,Student{"Mort Huang",18,"Earth"},http.StatusOK)
	}).All()
	log.Fatal(http.ListenAndServe(":8080", router))
}