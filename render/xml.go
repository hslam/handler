package render
import (
	"net/http"
	"encoding/xml"
	"hslam.com/mgit/Mort/mux-x/header"
)

func XML(w http.ResponseWriter, r *http.Request, v interface{}, code int) (int,error) {
	var (
		body []byte
		err error
	)
	if r.FormValue("pretty") != ""{
		body, err = xml.MarshalIndent(v, "", "  ")
	} else {
		body, err = xml.Marshal(v)
	}
	if err != nil {
		return 0,err
	}
	header.SetContentTypeWithUTF8(w,header.ContentTypeXML)
	return Body(w,r,body,code)
}
