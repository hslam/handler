# render
Package render supports rendering for the http.ResponseWriter.

## Features
* JSON
* XML
* Redirect
* Text
* File
* Body
* Gzip
* Deflate
* Template

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


#### HelloWorld Example
```
package main

import (
	"github.com/hslam/handler/render"
	"github.com/hslam/mux"
	"log"
	"net/http"
)

func main() {
	r := render.NewRender()
	m := mux.New()
	m.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.Text(w, req, "Hello world", http.StatusOK)
	}).All()
	log.Fatal(http.ListenAndServe(":8080", m))
}
```
curl http://localhost:8080/
```
Hello world
```


#### Gzip/Deflate Example
```
package main

import (
	"github.com/hslam/handler/render"
	"github.com/hslam/mux"
	"log"
	"net/http"
)

func main() {
	r:=render.NewRender()
	r.GzipAll().DeflateAll().Charset("utf-8")
	m := mux.New()
	m.HandleFunc("/compress", func(w http.ResponseWriter, req *http.Request) {
		r.Body(w,req,[]byte("compress"),http.StatusOK)
	}).GET().POST().HEAD()
	log.Fatal(http.ListenAndServe(":8080", m))
}
```
curl -H "Accept-Encoding: gzip" --compressed http://localhost:8080/compress
```
compress
```

curl -H "Accept-Encoding: gzip" -I  --compressed http://localhost:8080/compress
```
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Type: text/plain; charset=utf-8
Vary: Accept-Encoding
Date: Sun, 06 Oct 2019 12:50:34 GMT
Content-Length: 37
```

curl -H "Accept-Encoding: deflate" --compressed http://localhost:8080/compress
```
compress
```
curl -H "Accept-Encoding: deflate" -I  --compressed http://localhost:8080/compress
```
HTTP/1.1 200 OK
Content-Encoding: deflate
Content-Type: text/plain; charset=utf-8
Vary: Accept-Encoding
Date: Sun, 06 Oct 2019 12:51:23 GMT
Content-Length: 25
```



#### Simple Example
```
package main

import (
	"github.com/hslam/handler/render"
	"github.com/hslam/mux"
	"log"
	"net/http"
)

type Student struct {
	Name    string
	Age     int32
	Address string
}

func main() {
	r := render.NewRender()
	r.GzipAll().DeflateAll().Charset("utf-8")
	m := mux.New()
	m.HandleFunc("/text", func(w http.ResponseWriter, req *http.Request) {
		r.Text(w, req, "Hello world", http.StatusOK)
	}).All()
	m.HandleFunc("/raw", func(w http.ResponseWriter, req *http.Request) {
		r.Body(w, req, []byte("raw data"), http.StatusOK)
	}).All()
	m.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, req, Student{"Mort Huang", 18, "Earth"}, http.StatusOK)
	}).All()
	m.HandleFunc("/xml", func(w http.ResponseWriter, req *http.Request) {
		r.XML(w, req, Student{"Mort Huang", 18, "Earth"}, http.StatusOK)
	}).All()
	log.Fatal(http.ListenAndServe(":8080", m))
}
```

curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/raw
```
raw data
```

curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/json?pretty=y
```
{
  "Name": "Meng Huang",
  "Age": 18,
  "Address": "Earth"
}
```

curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/xml?pretty=y
```
<Student>
  <Name>Meng Huang</Name>
  <Age>18</Age>
  <Address>Earth</Address>
</Student>
```

#### Template Example
```
package main

import (
	"fmt"
	"github.com/hslam/handler/render"
	"github.com/hslam/mux"
	"log"
	"net/http"
)

type Student struct {
	Name    string
	Age     int32
	Address string
}

var studentTemplate = `
This is a Student Template:
Name: {{.Name}};
Age:{{.Age}};
Address: {{.Address}}.
`

var studentTemplateOne = `
<br>Name	-	{{.Name}}<br/>
<br>Age		-	{{.Age}}<br/>
<br>Address	-	{{.Address}}<br/>
`

var studentTemplateTwo = `
<html><body>
<table style='width:100%'>
	<tr><td align='center'>Name</td><td align='center'>{{.Name}}</td><tr>
	<tr><td align='center'>Age</td><td align='center'>{{.Age}}</td><tr>
	<tr><td align='center'>Address</td><td align='center'>{{.Address}}</td><tr>
</table>
</body></html>
`

func main() {
	r := render.NewRender()
	r.Parse(studentTemplate)
	r.ParseTemplate("1", studentTemplateOne)
	r.ParseTemplate("2", studentTemplateTwo)
	r.GzipAll().DeflateAll().Charset("utf-8")
	m := mux.New()
	m.HandleFunc("/template", func(w http.ResponseWriter, req *http.Request) {
		r.Execute(w, req, Student{"Mort Huang", 18, "Earth"}, http.StatusOK)
	}).All()
	m.HandleFunc("/template/:name", func(w http.ResponseWriter, req *http.Request) {
		params := m.Params(req)
		_, err := r.ExecuteTemplate(w, req, params["name"], Student{"Mort Huang", 18, "Earth"}, http.StatusOK)
		if err != nil {
			r.Text(w, req, fmt.Sprintf("template/%s is not exsited", params["name"]), http.StatusOK)
		}
	}).All()
	log.Fatal(http.ListenAndServe(":8080", m))
}
```
curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/template
```
This is a Student Template:
Name: Meng Huang;
Age:18;
Address: Earth.
```
curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/template/1
```
<br>Name	-	Meng Huang<br/>
<br>Age		-	18<br/>
<br>Address	-	Earth<br/>
```

curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/template/2
```
<html><body>
<table style='width:100%'>
	<tr><td align='center'>Name</td><td align='center'>Meng Huang</td><tr>
	<tr><td align='center'>Age</td><td align='center'>18</td><tr>
	<tr><td align='center'>Address</td><td align='center'>Earth</td><tr>
</table>
</body></html>
```

### License
This package is licensed under a MIT license (Copyright (c) 2019 Meng Huang)


### Authors
render was written by Meng Huang.


