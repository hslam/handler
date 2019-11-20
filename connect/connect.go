package connect

import (
	"net/http"
	"net"
	"io"
	"bufio"
	"errors"
	"net/url"
)

var connected = "200 Connected to Mux"

func GetConn(w http.ResponseWriter, r *http.Request)net.Conn{
	if r.Method != "CONNECT" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "405 must CONNECT\n")
		return nil
	}
	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		return nil
	}
	io.WriteString(conn, "HTTP/1.0 "+connected+"\n\n")
	return conn
}
func DialHTTP(u string)(net.Conn, error) {
	u_parse,err:=url.Parse(u)
	if err!=nil{
		return nil,err
	}
	return DialHTTPPath(u_parse.Host,u_parse.Path)
}
func DialHTTPPath(address, path string) (net.Conn, error) {
	var err error
	var network = "tcp"
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	io.WriteString(conn, "CONNECT "+path+" HTTP/1.0\n\n")

	// Require successful HTTP response
	// before switching to Your protocol.
	resp, err := http.ReadResponse(bufio.NewReader(conn), &http.Request{Method: "CONNECT"})
	if err == nil && resp.Status == connected {
		return conn, nil
	}
	if err == nil {
		err = errors.New("unexpected HTTP response: " + resp.Status)
	}
	conn.Close()
	return nil, &net.OpError{
		Op:   "dial-http",
		Net:  network + " " + address,
		Addr: nil,
		Err:  err,
	}
}