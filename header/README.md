# header
Package header sets header for the http.ResponseWriter.

## Get started

### Install
```
go get github.com/hslam/handler
```
### Import
```
import "github.com/hslam/handler"
```
### Usage
#### Example
```go
package main

import (
	"github.com/hslam/handler/header"
	"github.com/hslam/mux"
	"log"
	"net/http"
)

func main() {
	m := mux.New()
	m.Use(func(w http.ResponseWriter, r *http.Request) {
		header.SetHeader(w, header.AccessControlAllowOrigin, "*")
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

### License
This package is licensed under a MIT license (Copyright (c) 2019 Meng Huang)


### Author
header was written by Meng Huang.


