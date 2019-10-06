package main
import (
	"log"
	"net/http"
	"hslam.com/mgit/Mort/mux"
	"hslam.com/mgit/Mort/handler/compress"
)
func main() {
	router := mux.New()
	router.HandleFunc("/gzip", func(w http.ResponseWriter, r *http.Request) {
		compress.Gzip(w,r,[]byte("gzip"),http.StatusOK)
	}).GET().POST().HEAD()
	router.HandleFunc("/deflate", func(w http.ResponseWriter, r *http.Request) {
		compress.Deflate(w,r,[]byte("deflate"),http.StatusOK)
	}).GET().POST().HEAD()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
}