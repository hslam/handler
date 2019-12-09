package main
import (
	"log"
	"net/http"
	"github.com/hslam/mux"
	"github.com/hslam/handler/proxy"
)
func main() {
	go func() {
		m := mux.New()
		m.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello from 8081"))
		}).All()
		log.Fatal(http.ListenAndServe(":8081", m))
	}()
	m := mux.New()
	m.HandleFunc("/proxy", func(w http.ResponseWriter, r *http.Request) {
		proxy.Proxy(w,r,"http://localhost:8081/hello")
	}).All()
	log.Fatal(http.ListenAndServe(":8080", m))
}