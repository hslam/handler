# compress
Package compress supports gzip/deflate compressing for the http.ResponseWriter.

## Features

* Gzip
* Deflate

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
	"github.com/hslam/handler/compress"
	"github.com/hslam/mux"
	"log"
	"net/http"
)

func main() {
	m := mux.New()
	m.HandleFunc("/gzip", func(w http.ResponseWriter, r *http.Request) {
		compress.Gzip(w, r, []byte("gzip"), http.StatusOK)
	}).GET().POST().HEAD()
	m.HandleFunc("/deflate", func(w http.ResponseWriter, r *http.Request) {
		compress.Deflate(w, r, []byte("deflate"), http.StatusOK)
	}).GET().POST().HEAD()
	log.Fatal(http.ListenAndServe(":8080", m))
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

### License
This package is licensed under a MIT license (Copyright (c) 2019 Meng Huang)


### Author
compress was written by Meng Huang.


