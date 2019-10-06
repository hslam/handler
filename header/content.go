package header

import (
	"net/http"
	"strconv"
)

func SetContentLength(w http.ResponseWriter, length int){
	SetHeader(w,ContentLength,strconv.Itoa(length))
}

func SetContentType(w http.ResponseWriter, value string){
	SetHeader(w,ContentType,value)
}

func SetCharset(w http.ResponseWriter, charset string){
	if contentType:=GetResponseHeader(w,ContentType); contentType!="" {
		SetContentTypeWithCharset(w,contentType,charset)
	}
}

func SetContentTypeWithUTF8(w http.ResponseWriter, value string){
	SetContentTypeWithCharset(w,value,UTF8)
}

func SetContentTypeWithCharset(w http.ResponseWriter, value string,charset string){
	SetHeader(w,ContentType,value+Semicolon+CharsetPrefix+charset)
}
