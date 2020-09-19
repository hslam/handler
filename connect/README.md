# connect
Package connect switchs protocol from HTTP to TCP connection using CONNECT HTTP method.
## Get started

### Install
```
go get github.com/hslam/handler
```
### Import
```
import "github.com/hslam/handler"
```
### Usage
#### Example

**server**
```go
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
		conn := connect.GetConn(w, r)
		ServeConn(conn)
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
```

**Client**
```go
package main

import (
	"bufio"
	"fmt"
	"github.com/hslam/handler/connect"
	"io"
	"time"
)

func main() {
	//conn,err:=connect.DialHTTP("http://localhost:8080/connect")
	conn, err := connect.DialHTTPPath("localhost:8080", "/connect")
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(conn)
	for i := 0; i < 3; i++ {
		conn.Write([]byte("Hello mux\n"))
		message, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Print(message)
		time.Sleep(time.Second)
	}
}
```

**Output**
```
HELLO MUX
HELLO MUX
HELLO MUX
```
### License
This package is licensed under a MIT license (Copyright (c) 2019 Meng Huang)


### Authors
connect was written by Meng Huang.


