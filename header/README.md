# header
## mux middleware to set header.


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
	"hslam.com/git/x/handler/header"
)
func main() {
	m := mux.New()
	m.Use(func(w http.ResponseWriter, r *http.Request) {
		header.SetHeader(w,header.AccessControlAllowOrigin, "*")
	})
	m.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).All()
	log.Fatal(http.ListenAndServe(":8080", m))
}
```
curl -I http://localhost:8080/hello
```
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Date: Sun, 06 Oct 2019 15:58:02 GMT
Content-Length: 11
Content-Type: text/plain; charset=utf-8
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)


### Authors
header was written by Mort Huang.


