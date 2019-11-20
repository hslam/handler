# connect
Switch from HTTP to TCP connection using CONNECT HTTP method.
## Get started

### Install
```
go get hslam.com/git/x/handler
```
### Import
```
import "hslam.com/git/x/handler"
```
### Usage
#### Example

**server**
```go
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
```

**Client**
```go
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
```

**Output**
```
Hello mux
Hello mux
Hello mux
...
```
### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Mort Huang)


### Authors
connect was written by Mort Huang.


