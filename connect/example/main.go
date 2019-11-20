package main
import (
	"log"
	"net/http"
	"hslam.com/git/x/mux"
	"hslam.com/git/x/handler/connect"
	"net"
	"bufio"
	"io"
)
func main() {
	router := mux.New()
	router.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		conn:=connect.GetConn(w,r)
		ServeConn(conn)
	}).CONNECT()
	router.Once()//before listening
	log.Fatal(http.ListenAndServe(":8080", router))
}
//Echo
func ServeConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for{
		message,err := reader.ReadString('\n')
		if err!=nil || err == io.EOF {
			break
		}
		conn.Write([]byte(message))
	}
}