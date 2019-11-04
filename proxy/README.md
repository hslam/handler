# proxy
## mux middleware to proxy to other server.

## Get started

### Install
```
go get hslam.com/git/x/handler
```
### Import
```
import "hslam.com/git/x/handler"
```
### Usage
#### Example
```
package main
import (
	"log"
	"net/http"
	"hslam.com/git/x/mux"
	"hslam.com/git/x/handler/proxy"
)
func main() {
	go func() {
		router := mux.New()
		router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello from 8081"))
		}).All()
		router.Once()//before listening
		log.Fatal(http.ListenAndServe(":8081", router))
	}()
	router := mux.New()
	router.HandleFunc("/proxy", func(w http.ResponseWriter, r *http.Request) {
		proxy.Proxy(w,r,"http://localhost:8081/hello")
	}).All()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
}
```
curl http://localhost:8080/proxy
```
hello from 8081
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)


### Authors
proxy was written by Mort Huang.


