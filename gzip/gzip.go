package gzip

import (
	"compress/gzip"
	"net/http"
	"strings"
	"bytes"
	"hslam.com/mgit/Mort/mux-x/header"
)

type GzipWriter struct {
	Writer *gzip.Writer
	w http.ResponseWriter
	gzip bool
	buf *bytes.Buffer
}

func NewGzipWriter(w http.ResponseWriter, r *http.Request)*GzipWriter  {
	g:=&GzipWriter{
		Writer:gzip.NewWriter(w),
		w:w,
	}
	g.ready(w,r)
	return g
}

func (g *GzipWriter)ready(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(header.GetRequestHeader(r,header.AcceptEncoding), header.Gzip) {
		g.gzip=false
		return
	}
	header.SetHeader(w,header.ContentEncoding,header.Gzip)
	header.SetHeader(w,header.Vary,header.AcceptEncoding)
	header.DelHeader(w,header.ContentLength)
	g.gzip=true
}

func (g *GzipWriter) Write(b []byte) (int, error) {
	if g.gzip{
		header.SetHeader(g.w,header.ContentType,http.DetectContentType(b))
		return g.Writer.Write(b)
	}else {
		return g.w.Write(b)
	}
}

func (g *GzipWriter) Close() (error) {
	if g.gzip{
		return g.Writer.Close()
	}
	return nil
}

func Gzip(w http.ResponseWriter, r *http.Request, body []byte,code int) (int, error) {
	w.WriteHeader(code)
	gz:=NewGzipWriter(w,r)
	defer gz.Close()
	return gz.Write(body)
}