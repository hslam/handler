package main
import (
	"log"
	"net/http"
	"hslam.com/mgit/Mort/mux"
	"hslam.com/mgit/Mort/handler/header"
)
func main() {
	router := mux.New()
	router.Use(func(w http.ResponseWriter, r *http.Request) {
		header.SetHeader(w,header.AccessControlAllowOrigin, "*")
	})
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).All()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
}