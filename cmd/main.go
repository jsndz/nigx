package main

import (
	"fmt"
	"net"
	"nigx/internals/http"
	"nigx/internals/static"
)

func HandlerRequests(conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(buf[:n]))
	req := http.NewHttpRequest(string(buf[:n]))
	var res *http.HttpResponse
	data, err := static.TryFiles(req.Url)
	if err != nil {
		res = http.NewHttpResponse("HTTP/1.1", 404, "Not Found", map[string]string{"Content-Type": "text/plain"}, []byte("404 Not Found"))
	} else {
		res = http.NewHttpResponse("HTTP/1.1", 200, "OK", map[string]string{"Content-Type": "text/html"}, data)
	}
	conn.Write(res.Bytes())
	conn.Close()
}
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
		go HandlerRequests(conn)
	}
}
