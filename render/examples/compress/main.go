package main
import (
	"log"
	"net/http"
	"hslam.com/git/x/rum"
	"hslam.com/git/x/handler/render"
)
func main() {
	r:=render.NewRender()
	r.GzipAll().DeflateAll().Charset("utf-8")
	router := rum.New()
	router.HandleFunc("/compress", func(w http.ResponseWriter, req *http.Request) {
		r.Body(w,req,[]byte("compress"),http.StatusOK)
	}).GET().POST().HEAD()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
}