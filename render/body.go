package render

import (
	"net/http"
	"hslam.com/mgit/Mort/mux-x/header"
)
func Body(w http.ResponseWriter, r *http.Request,body []byte, code int)(int,error)  {
	header.SetHeaderLength(w,len(body))
	header.SetHeader(w,header.ContentType,http.DetectContentType(body))
	w.WriteHeader(code)
	return w.Write(body)
}