package render

import (
	"net/http"
	"io/ioutil"
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
	return Body(w,r,body,code)
}

func loadFile(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}

func ServeFile(w http.ResponseWriter, r *http.Request,name string) {
	http.ServeFile(w,r,name)
}