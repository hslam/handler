package main
import (
	"log"
	"net/http"
	"hslam.com/git/x/mux"
	"hslam.com/git/x/handler/render"
)
func main() {
	r:=render.NewRender()
	router := mux.New()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.Text(w,req,"Hello world",http.StatusOK)
	}).All()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
}