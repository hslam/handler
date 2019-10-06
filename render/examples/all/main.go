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
	r.GzipAll().DeflateAll().Charset("utf-8")
	r.Parse(studentTemplate)
	r.ParseTemplate("1",studentTemplateOne)
	r.ParseTemplate("2",studentTemplateTwo)
	router := mux.New()
	router.HandleFunc("/compress", func(w http.ResponseWriter, req *http.Request) {
		r.Body(w,req,[]byte("compress"),http.StatusOK)
	}).GET().POST().HEAD()
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
	router.HandleFunc("/template", func(w http.ResponseWriter, req *http.Request) {
		r.Execute(w,req,Student{"Mort Huang",18,"Earth"},http.StatusOK)
	}).All()
	router.HandleFunc("/template/:name", func(w http.ResponseWriter, req *http.Request) {
		params:=router.Params(req)
		if params["name"]!=""{
			_,err:=r.ExecuteTemplate(w,req,params["name"],Student{"Mort Huang",18,"Earth"},http.StatusOK)
			if err!=nil{
				r.Text(w,req,fmt.Sprintf("template/%s is not exsited",params["name"]),http.StatusOK)
			}
		}else {
			r.Text(w,req,"name is empty",http.StatusOK)
		}
	}).All()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
}