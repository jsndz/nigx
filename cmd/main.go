package main

import (
	"fmt"
	"net"
	"nigx/internals/http"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	conn, err := lis.Accept()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	fmt.Print(string(buf[:n]))
	http.NewHttpRequest(string(buf[:n]))
}
