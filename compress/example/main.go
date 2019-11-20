package main
import (
	"log"
	"net/http"
	"hslam.com/git/x/mux"
	"hslam.com/git/x/handler/compress"
)
func main() {
	m := mux.New()
	m.HandleFunc("/gzip", func(w http.ResponseWriter, r *http.Request) {
		compress.Gzip(w,r,[]byte("gzip"),http.StatusOK)
	}).GET().POST().HEAD()
	m.HandleFunc("/deflate", func(w http.ResponseWriter, r *http.Request) {
		compress.Deflate(w,r,[]byte("deflate"),http.StatusOK)
	}).GET().POST().HEAD()
	log.Fatal(http.ListenAndServe(":8080", m))
}