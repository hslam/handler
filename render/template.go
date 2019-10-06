package render

import (
	"sync"
	"bytes"
	"net/http"
	"html/template"
	"hslam.com/mgit/Mort/handler/header"
	"errors"
)

const defaultTemplateName  = "Html"

var bufPool *sync.Pool

func init() {
	bufPool=&sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
}

type Tmpl struct {
	t *template.Template
}

func NewTmpl() *Tmpl{
	return &Tmpl{}
}

func (t *Tmpl)NewTemplate(text string)(error){
	tmpl, err := template.New(defaultTemplateName).Parse(text)
	if err!=nil{
		return err
	}
	t.t=tmpl
	return nil
}

func (t *Tmpl)Execute(w http.ResponseWriter, r *http.Request,data interface{}, code int) (int,error){
	if t.t==nil{
		return 0,errors.New("Must NewTemplate")
	}
	html,err:=t.execute(data)
	if err!=nil{
		return 0,err
	}
	header.SetContentTypeWithUTF8(w,header.ContentTypeHTML)
	return Body(w,r,[]byte(html),code)
}

func (t *Tmpl)execute(data interface{}) (string,error) {
	buf := bufPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		bufPool.Put(buf)
	}()
	if err:=t.t.Execute(buf, data);err!=nil{
		return "",err
	}
	return buf.String(),nil
}

func (t *Tmpl)ParseFiles(filenames ...string)(error){
	tmpl, err := template.ParseFiles(filenames...)
	if err!=nil{
		return err
	}
	t.t=tmpl
	return nil
}
func (t *Tmpl)ExecuteTemplate(w http.ResponseWriter, r *http.Request,name string,data interface{}, code int) (int,error){
	if t.t==nil{
		return 0,errors.New("Must ParseFiles")
	}
	html,err:=t.executeTemplate(name,data)
	if err!=nil{
		return 0,err
	}
	header.SetContentTypeWithUTF8(w,header.ContentTypeHTML)
	return Body(w,r,[]byte(html),code)
}
func (t *Tmpl)executeTemplate(name string,data interface{}) (string,error) {
	buf := bufPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		bufPool.Put(buf)
	}()
	if err:=t.t.ExecuteTemplate(buf, name,data);err!=nil{
		return "",err
	}
	return buf.String(),nil
}

func Execute(w http.ResponseWriter, r *http.Request, text string,data interface{}, code int) (int,error){
	t:=NewTmpl()
	err:=t.NewTemplate(text)
	if err!=nil{
		return 0,err
	}
	return t.Execute(w,r,data,code)
}

func ExecuteTemplate(w http.ResponseWriter, r *http.Request, filename,name string,data interface{}, code int) (int,error){
	t:=NewTmpl()
	err:=t.ParseFiles(filename)
	if err!=nil{
		return 0,err
	}
	return t.ExecuteTemplate(w,r,name,data,code)
}
