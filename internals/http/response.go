package http

type HttpResponse struct {
	FullResponse []byte
	// for now
	// return full http response
}

func NewHttpResponse() []byte {
	return []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 12\r\n\r\nHello World!")
}
