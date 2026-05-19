package http

import (
	"bytes"
	"fmt"
)

type HttpResponse struct {
	Version    string
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       []byte
}

func NewHttpResponse(version string, statusCode int, statusText string, headers map[string]string, body []byte) *HttpResponse {
	return &HttpResponse{
		Version:    version,
		StatusCode: statusCode,
		StatusText: statusText,
		Headers:    headers,
		Body:       body,
	}
}

func (res *HttpResponse) Bytes() []byte {
	var buf bytes.Buffer

	fmt.Fprintf(
		&buf,
		"%s %d %s\r\n",
		res.Version,
		res.StatusCode,
		res.StatusText,
	)

	for key, value := range res.Headers {
		fmt.Fprintf(&buf, "%s: %s\r\n", key, value)
	}

	buf.WriteString("\r\n")

	buf.Write(res.Body)

	return buf.Bytes()
}
