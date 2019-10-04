package host

import (
	"strings"
	"net/http"
)
func Host(req *http.Request) (addr string) {
	return ParseHostName(req.Referer())
}
func ParseHostName(url string) string {
	url = strings.TrimRight(url, "/")
	url = strings.TrimRight(url, "/?")
	if len(url) == 0 {
		return ""
	}
	strs := strings.Split(url, "//")
	if len(strs) == 0 {
		return ""
	}else if len(strs) == 1 {
		return ""
	}
	host_name := strs[1]
	strs = strings.Split(host_name, "/")
	if len(strs) == 0 {
		return ""
	}
	return strs[0]
}