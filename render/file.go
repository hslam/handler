package render

import (
	"net/http"
	"io/ioutil"
	"hslam.com/mgit/Mort/mux-x/header"
)

func File(w http.ResponseWriter, r *http.Request,name string,code int) (int,error) {
	var (
		body []byte
		err error
	)
	body, err = loadFile(name)
	if err!=nil{
		return 0,err
	}
	header.SetHeader(w,header.ContentType,http.DetectContentType(body))
	w.WriteHeader(code)
	return w.Write(body)
}

func loadFile(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}

func ServeFile(w http.ResponseWriter, r *http.Request,name string) {
	http.ServeFile(w,r,name)
}