package render

import (
	"net/http"
	"hslam.com/mgit/Mort/handler/header"
)
func Text(w http.ResponseWriter, r *http.Request,text string, code int)(int,error)  {
	header.SetContentTypeWithUTF8(w,header.ContentTypeText)
	return Body(w,r,[]byte(text),code)
}