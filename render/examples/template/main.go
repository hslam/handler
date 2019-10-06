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