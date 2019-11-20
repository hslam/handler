package main
import (
	"log"
	"net/http"
	"hslam.com/git/x/mux"
	"hslam.com/git/x/handler/header"
)
func main() {
	m := mux.New()
	m.Use(func(w http.ResponseWriter, r *http.Request) {
		header.SetHeader(w,header.AccessControlAllowOrigin, "*")
	})
	m.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).All()
	log.Fatal(http.ListenAndServe(":8080", m))
}