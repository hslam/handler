package render

import (
	"net/http"
	"hslam.com/mgit/Mort/mux-x/header"
)
func Text(w http.ResponseWriter, r *http.Request,text string, code int)(int,error)  {
	header.SetHeaderLength(w,len(text))
	header.SetContentTypeUTF8(w,header.ContentTypeText)
	w.WriteHeader(code)
	return w.Write([]byte(text))
}