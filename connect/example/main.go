package main

import (
	"bufio"
	"github.com/hslam/handler/connect"
	"github.com/hslam/mux"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	m := mux.New()
	m.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		if conn, err := connect.Connect(w, r); err == nil {
			ServeConn(conn)
		}
	}).CONNECT()
	log.Fatal(http.ListenAndServe(":8080", m))
}

func ServeConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		conn.Write([]byte(strings.ToUpper(string(message))))
	}
}
