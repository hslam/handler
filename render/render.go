package render

import (
	"encoding/json"
	"encoding/xml"
	"github.com/hslam/handler/compress"
	"github.com/hslam/handler/header"
	"io/ioutil"
	"net/http"
	"sync"
)

var DefalutRender *Render

func init() {
	DefalutRender = NewRender()
}

type Render struct {
	mut            sync.RWMutex
	charset        string
	gzip           bool
	deflate        bool
	compressWriter *compress.CompressWriter
	tmpl           *Tmpl
}

func NewRender() *Render {
	render := &Render{charset: header.UTF8}
	render.tmpl = NewTmplWithRender(render)
	return render
}

func (render *Render) GzipAll() *Render {
	render.mut.Lock()
	defer render.mut.Unlock()
	render.gzip = true
	return render
}

func (render *Render) DeflateAll() *Render {
	render.mut.Lock()
	defer render.mut.Unlock()
	render.deflate = true
	return render
}

func (render *Render) Charset(charset string) *Render {
	render.mut.Lock()
	defer render.mut.Unlock()
	render.charset = charset
	return render
}

func Body(w http.ResponseWriter, r *http.Request, body []byte, code int) (int, error) {
	return DefalutRender.Body(w, r, body, code)
}

func (render *Render) write(w http.ResponseWriter, r *http.Request, body []byte, code int) (int, error) {
	render.mut.RLock()
	defer render.mut.RUnlock()
	header.SetContentLength(w, len(body))
	if contentType := header.GetResponseHeader(w, header.ContentType); contentType == "" {
		header.SetHeader(w, header.ContentType, http.DetectContentType(body))
	}
	if render.deflate && header.CheckAcceptEncoding(r, header.DEFLATE) {
		c := compress.NewDeflateWriter(w, r)
		w.WriteHeader(code)
		n, err := c.Write(body)
		defer c.Close()
		return n, err
	} else if render.gzip && header.CheckAcceptEncoding(r, header.GZIP) {
		c := compress.NewGzipWriter(w, r)
		w.WriteHeader(code)
		n, err := c.Write(body)
		defer c.Close()
		return n, err
	}
	w.WriteHeader(code)
	return w.Write(body)
}

func (render *Render) Body(w http.ResponseWriter, r *http.Request, body []byte, code int) (int, error) {
	return render.write(w, r, body, code)
}

func File(w http.ResponseWriter, r *http.Request, name string, code int) (int, error) {
	return DefalutRender.File(w, r, name, code)
}

func (render *Render) File(w http.ResponseWriter, r *http.Request, name string, code int) (int, error) {
	var (
		body []byte
		err  error
	)
	body, err = ioutil.ReadFile(name)
	if err != nil {
		return 0, err
	}
	return render.write(w, r, body, code)
}

func ServeFile(w http.ResponseWriter, r *http.Request, name string) {
	DefalutRender.ServeFile(w, r, name)
}

func (render *Render) ServeFile(w http.ResponseWriter, r *http.Request, name string) {
	http.ServeFile(w, r, name)
}

func JSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) (int, error) {
	return DefalutRender.JSON(w, r, v, code)
}

func (render *Render) JSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) (int, error) {
	var (
		body []byte
		err  error
	)
	if r.FormValue("pretty") != "" {
		body, err = json.MarshalIndent(v, "", "  ")
	} else {
		body, err = json.Marshal(v)
	}
	if err != nil {
		return 0, err
	}
	header.SetContentTypeWithCharset(w, header.ContentTypeJSON, render.charset)
	return render.write(w, r, body, code)
}

func XML(w http.ResponseWriter, r *http.Request, v interface{}, code int) (int, error) {
	return DefalutRender.XML(w, r, v, code)
}

func (render *Render) XML(w http.ResponseWriter, r *http.Request, v interface{}, code int) (int, error) {
	var (
		body []byte
		err  error
	)
	if r.FormValue("pretty") != "" {
		body, err = xml.MarshalIndent(v, "", "  ")
	} else {
		body, err = xml.Marshal(v)
	}
	if err != nil {
		return 0, err
	}
	header.SetContentTypeWithCharset(w, header.ContentTypeXML, render.charset)
	return render.write(w, r, body, code)
}

func Redirect(w http.ResponseWriter, r *http.Request, url string) {
	DefalutRender.Redirect(w, r, url)
}

func (render *Render) Redirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusFound)
}

func Text(w http.ResponseWriter, r *http.Request, text string, code int) (int, error) {
	return DefalutRender.Text(w, r, text, code)
}

func (render *Render) Text(w http.ResponseWriter, r *http.Request, text string, code int) (int, error) {
	header.SetContentTypeWithCharset(w, header.ContentTypeText, render.charset)
	return render.write(w, r, []byte(text), code)
}

func (render *Render) Parse(text string) error {
	return render.tmpl.Parse(text)
}

func (render *Render) Execute(w http.ResponseWriter, r *http.Request, data interface{}, code int) (int, error) {
	return render.tmpl.Execute(w, r, data, code)
}

func (render *Render) ParseTemplate(name, text string) error {
	return render.tmpl.ParseTemplate(name, text)
}

func (render *Render) ExecuteTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}, code int) (int, error) {
	return render.tmpl.ExecuteTemplate(w, r, name, data, code)
}
