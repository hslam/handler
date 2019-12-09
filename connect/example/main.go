package main
import (
	"log"
	"net/http"
	"github.com/hslam/mux"
	"github.com/hslam/handler/connect"
	"net"
	"bufio"
	"io"
	"strings"
)
func main() {
	m := mux.New()
	m.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		conn:=connect.GetConn(w,r)
		ServeConn(conn)
	}).CONNECT()
	log.Fatal(http.ListenAndServe(":8080", m))
}
//
func ServeConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for{
		message,err := reader.ReadString('\n')
		if err!=nil || err == io.EOF {
			break
		}
		conn.Write([]byte(strings.ToUpper(string(message))))
	}
}