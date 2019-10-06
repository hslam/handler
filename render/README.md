# render
## mux middleware to enable render support.

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
go get hslam.com/mgit/Mort/handler
```
### Import
```
import "hslam.com/mgit/Mort/handler"
```
### Usage


#### HelloWorld Example
```
package main
import (
	"log"
	"net/http"
	"hslam.com/mgit/Mort/mux"
	"hslam.com/mgit/Mort/handler/render"
)
func main() {
	r:=render.NewRender()
	router := mux.New()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.Text(w,req,"Hello world",http.StatusOK)
	}).All()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
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
	"log"
	"net/http"
	"hslam.com/mgit/Mort/mux"
	"hslam.com/mgit/Mort/handler/render"
)
func main() {
	ren:=render.NewRender()
	ren.GzipAll().DeflateAll()
	router := mux.New()
	router.HandleFunc("/compress", func(w http.ResponseWriter, r *http.Request) {
		ren.Body(w,r,[]byte("compress"),http.StatusOK)
	}).GET().POST().HEAD()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
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
	"log"
	"net/http"
	"hslam.com/mgit/Mort/mux"
	"hslam.com/mgit/Mort/handler/render"
)
type Student struct {
	Name string
	Age int32
	Address string
}
func main() {
	r:=render.NewRender()
	r.GzipAll().DeflateAll()
	router := mux.New()
	router.HandleFunc("/text", func(w http.ResponseWriter, req *http.Request) {
		r.Text(w,req,"Hello world",http.StatusOK)
	}).All()
	router.HandleFunc("/raw", func(w http.ResponseWriter, req *http.Request) {
		r.Body(w,req,[]byte("raw data"),http.StatusOK)
	}).All()
	router.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w,req,Student{"Mort Huang",18,"Earth"},http.StatusOK)
	}).All()
	router.HandleFunc("/xml", func(w http.ResponseWriter, req *http.Request) {
		r.XML(w,req,Student{"Mort Huang",18,"Earth"},http.StatusOK)
	}).All()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
}
```

curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/raw
```
raw data
```

curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/json?pretty=y
```
{
  "Name": "Mort Huang",
  "Age": 18,
  "Address": "Earth"
}
```

curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/xml?pretty=y
```
<Student>
  <Name>Mort Huang</Name>
  <Age>18</Age>
  <Address>Earth</Address>
</Student>
```

#### Template Example
```
package main
import (
	"log"
	"net/http"
	"hslam.com/mgit/Mort/mux"
	"hslam.com/mgit/Mort/handler/render"
	"fmt"
)
type Student struct {
	Name string
	Age int32
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
	r:=render.NewRender()
	r.Parse(studentTemplate)
	r.ParseTemplate("1",studentTemplateOne)
	r.ParseTemplate("2",studentTemplateTwo)
	r.GzipAll().DeflateAll()
	router := mux.New()
	router.HandleFunc("/template", func(w http.ResponseWriter, req *http.Request) {
		r.Execute(w,req,Student{"Mort Huang",18,"Earth"},http.StatusOK)
	}).All()
	router.HandleFunc("/template/:name", func(w http.ResponseWriter, req *http.Request) {
		params:=router.Params(req)
		_,err:=r.ExecuteTemplate(w,req,params["name"],Student{"Mort Huang",18,"Earth"},http.StatusOK)
		if err!=nil{
			r.Text(w,req,fmt.Sprintf("template/%s is not exsited",params["name"]),http.StatusOK)
		}
	}).All()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
}
```
curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/template
```
This is a Student Template:
Name: Mort Huang;
Age:18;
Address: Earth.
```
curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/template/1
```
<br>Name	-	Mort Huang<br/>
<br>Age		-	18<br/>
<br>Address	-	Earth<br/>
```

curl -H "Accept-Encoding: gzip,deflate" --compressed http://localhost:8080/template/2
```
<html><body>
<table style='width:100%'>
	<tr><td align='center'>Name</td><td align='center'>Mort Huang</td><tr>
	<tr><td align='center'>Age</td><td align='center'>18</td><tr>
	<tr><td align='center'>Address</td><td align='center'>Earth</td><tr>
</table>
</body></html>
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)


### Authors
render was written by Mort Huang.


