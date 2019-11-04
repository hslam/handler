# compress
## mux middleware to enable gzip and deflate support.

## Features

* Gzip
* Deflate

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
	"hslam.com/git/x/handler/compress"
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
```
curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/gzip
```
gzip
```

curl -H "Accept-Encoding: gzip,deflate" -I  --compressed http://localhost:8080/gzip
```
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Type: text/plain; charset=utf-8
Vary: Accept-Encoding
Date: Sun, 06 Oct 2019 11:43:10 GMT
Content-Length: 20
```

curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/deflate
```
deflate
```
curl -H "Accept-Encoding: gzip,deflate" -I  --compressed http://localhost:8080/deflate
```
HTTP/1.1 200 OK
Content-Encoding: deflate
Content-Type: text/plain; charset=utf-8
Vary: Accept-Encoding
Date: Sun, 06 Oct 2019 11:43:28 GMT
Content-Length: 15
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)


### Authors
mux was written by Mort Huang.


