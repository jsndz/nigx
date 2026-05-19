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

	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		go func() {

			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				panic(err)
			}
			fmt.Print(string(buf[:n]))
			http.NewHttpRequest(string(buf[:n]))
			conn.Write(http.NewHttpResponse())
		}()
	}
}
