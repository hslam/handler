package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

const MaxConnsPerHost = 16384

var transport *http.Transport

func init() {
	transport = &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		MaxIdleConnsPerHost: MaxConnsPerHost,
	}
}

func Proxy(w http.ResponseWriter, r *http.Request, targetUrl string) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	targetUrlParse, err := url.Parse(targetUrl)
	if err != nil {
		panic(err)
	}
	target, err := url.Parse("http://" + targetUrlParse.Host)
	if err != nil {
		panic(err)
	}
	r.URL.Path = targetUrlParse.Path
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = transport
	proxy.ServeHTTP(w, r)
}
