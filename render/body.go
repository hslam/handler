package render

import (
	"net/http"
	"hslam.com/mgit/Mort/handler/header"
)
func Body(w http.ResponseWriter, r *http.Request,body []byte, code int)(int,error)  {
	header.SetContentLength(w,len(body))
	if contentType:= header.GetResponseHeader(w,header.ContentType); contentType=="" {
		header.SetHeader(w,header.ContentType,http.DetectContentType(body))
	}
	w.WriteHeader(code)
	return w.Write(body)
}