package main
import (
	"log"
	"net/http"
	"hslam.com/mgit/Mort/mux"
	"hslam.com/mgit/Mort/handler/render"
)
func main() {
	ren:=render.NewRender()
	ren.GzipAll().DeflateAll()
	router := mux.New()
	router.HandleFunc("/compress", func(w http.ResponseWriter, r *http.Request) {
		ren.Body(w,r,[]byte("compress"),http.StatusOK)
	}).GET().POST().HEAD()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
}