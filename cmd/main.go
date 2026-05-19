package main

import (
	"net"
	"nigx/internals/config"
	"nigx/internals/http"
	"nigx/internals/proxy"
	"nigx/internals/static"
)

func HandlerRequests(conn net.Conn, cfg *config.Config) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	req := http.NewHttpRequest(string(buf[:n]))
	var res *http.HttpResponse

	if static.TryFiles(req.Url) {
		data, err := static.GetStaticFiles(req.Url)
		if err != nil {
			res = http.NewHttpResponse("HTTP/1.1", 404, "Not Found", map[string]string{"Content-Type": "text/plain"}, []byte("404 Not Found"))
		} else {
			res = http.NewHttpResponse("HTTP/1.1", 200, "OK", map[string]string{"Content-Type": "text/html"}, data)
		}
		conn.Write(res.Bytes())

	}
	if params, ok := proxy.IsProxyRequest(cfg.Route, req.Url); ok {
		respByte := proxy.ProxyRequest(req, cfg.Proxy, params)
		conn.Write(respByte)

	}
	conn.Close()

}
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	cfg := config.NewConfig("/api/", "https://jsonplaceholder.typicode.com/")
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		go HandlerRequests(conn, cfg)
	}
}
