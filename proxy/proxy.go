package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

const MaxConnsPerHost = 65536

var http_transport *http.Transport

func init() {
	http_transport = &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		MaxIdleConnsPerHost: MaxConnsPerHost,
		MaxConnsPerHost:     MaxConnsPerHost,
	}
}

func Proxy(w http.ResponseWriter, r *http.Request, target_url string) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	target_url_parse, err := url.Parse(target_url)
	if err != nil {
		panic(err)
	}
	target, err := url.Parse("http://" + target_url_parse.Host)
	if err != nil {
		panic(err)
	}
	r.URL.Path = target_url_parse.Path
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = http_transport
	proxy.ServeHTTP(w, r)
}
