package main
import (
	"github.com/hslam/handler/connect"
	"bufio"
	"io"
	"fmt"
	"time"
)
func main() {
	//conn,err:=connect.DialHTTP("http://localhost:8080/connect")
	conn,err:=connect.DialHTTPPath("localhost:8080","/connect")
	defer conn.Close()
	if err!=nil{
		panic(err)
	}
	reader := bufio.NewReader(conn)
	for i:=0;i<3;i++{
		conn.Write([]byte("Hello mux\n"))
		message,err := reader.ReadString('\n')
		if err!=nil || err == io.EOF {
			break
		}
		fmt.Print(message)
		time.Sleep(time.Second)
	}
}
