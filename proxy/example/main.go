package main
import (
	"log"
	"net/http"
	"hslam.com/git/x/rum"
	"hslam.com/git/x/handler/proxy"
)
func main() {
	go func() {
		router := rum.New()
		router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello from 8081"))
		}).All()
		router.Once()//before listening
		log.Fatal(http.ListenAndServe(":8081", router))
	}()
	router := rum.New()
	router.HandleFunc("/proxy", func(w http.ResponseWriter, r *http.Request) {
		proxy.Proxy(w,r,"http://localhost:8081/hello")
	}).All()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
}