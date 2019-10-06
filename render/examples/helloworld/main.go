package main
import (
	"log"
	"net/http"
	"hslam.com/mgit/Mort/mux"
	"hslam.com/mgit/Mort/handler/render"
)
func main() {
	r:=render.NewRender()
	router := mux.New()
	router.HandleFunc("/text", func(w http.ResponseWriter, req *http.Request) {
		r.Text(w,req,"Hello wolrd",http.StatusOK)
	}).All()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
}