package main
import (
	"hslam.com/git/x/handler/connect"
	"bufio"
	"io"
	"fmt"
	"time"
)
func main() {
	conn,err:=connect.DialHTTPPath("localhost:8080","/connect")
	if err!=nil{
		panic(err)
	}
	reader := bufio.NewReader(conn)
	for{
		conn.Write([]byte("Hello mux\n"))
		message,err := reader.ReadString('\n')
		if err!=nil || err == io.EOF {
			break
		}
		fmt.Print(message)
		time.Sleep(time.Second)
	}
}
