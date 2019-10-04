package header

import (
	"net/http"
	"strings"
	"strconv"
)
const (
	ContentEncoding		= "Content-Encoding"
	ContentLanguage		= "Content-Language"
	ContentLength  		= "Content-Length"
	ContentLocation  	= "Content-Location"
	ContentMD5  		= "Content-MD5"
	ContentRange  		= "Content-Range"
	ContentType    		= "Content-Type"

	AcceptRanges		= "Accept-Ranges"
	Bytes				= "bytes"

	AcceptEncoding		= "Accept-Encoding"
	Gzip 				= "gzip"

	Vary    			= "Vary"
	Charset 			= "charset"
	CharsetPrefix 		= "charset="
	UTF8 				= "UTF-8"
	GB18030 			= "GB18030"
	GBK 				= "GBK"

	Semicolon			= ";"
	Comma				= ","

	ContentTypeJSON    	= "application/json"
	ContentTypeXML     	= "text/xml"
	ContentTypeHTML    	= "text/html"
	ContentTypeText    	= "text/plain"
)
func SetHeader(w http.ResponseWriter,key, value string){
	if _, ok := w.Header()[key]; ok {
		w.Header().Set(key, value)
	} else {
		w.Header().Add(key, value)
	}
}
func DelHeader(w http.ResponseWriter,key string){
	w.Header().Del(key)
}

func WriteHeader(w http.ResponseWriter,code int){
	w.WriteHeader(code)
}
func SetHeaderLength(w http.ResponseWriter, length int){
	SetHeader(w,ContentLength,strconv.Itoa(length))
}
func SetContentTypeUTF8(w http.ResponseWriter, value string){
	SetHeader(w,ContentType,value+Semicolon+CharsetPrefix+UTF8)
}
func SetCharset(w http.ResponseWriter, charset string){
	if contentType, ok := w.Header()[ContentType]; ok {
		w.Header().Set(ContentType, strings.Join(contentType,Semicolon)+CharsetPrefix+charset)
	}
}

func GetHeader(r *http.Request,key string)(value string){
	return r.Header.Get(key)
}