package render

import (
	"net/http"
	"encoding/json"
	"hslam.com/mgit/Mort/mux-x/header"
)

func JSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) (int,error) {
	var (
		body []byte
		err error
	)
	if r.FormValue("pretty") != ""{
		body, err = json.MarshalIndent(v, "", "  ")
	} else {
		body, err = json.Marshal(v)
	}
	if err != nil {
		return 0,err
	}
	header.SetContentTypeUTF8(w,header.ContentTypeJSON)
	w.WriteHeader(code)
	return w.Write(body)
}
