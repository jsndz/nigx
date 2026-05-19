package http

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type HttpMethod string

const (
	POST   HttpMethod = "POST"
	GET    HttpMethod = "GET"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

type HttpRequest struct {
	Method        HttpMethod
	Version       int
	Url           string
	Host          string
	UserAgent     string
	ContentType   string
	ContentLength int
	Body          string
}

func NewHttpRequest(request string) *HttpRequest {
	lines := strings.Split(request, "\r\n")

	if len(lines) == 0 {
		return nil
	}

	requestLine := strings.Split(lines[0], " ")
	if len(requestLine) < 3 {
		return nil
	}

	method := HttpMethod(requestLine[0])

	version := 1
	if strings.Contains(requestLine[2], "1.0") {
		version = 0
	}

	req := &HttpRequest{
		Method:  method,
		Version: version,
		Url:     requestLine[1],
	}

	for i := 1; i < len(lines); i++ {
		line := lines[i]

		if line == "" {
			if i+1 < len(lines) {
				body := strings.Join(lines[i+1:], "\r\n")
				req.Body = body
			}
			break
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch strings.ToLower(key) {
		case "host":
			req.Host = value

		case "user-agent":
			req.UserAgent = value

		case "content-type":
			req.ContentType = value

		case "content-length":
			contentLength, err := strconv.Atoi(value)
			if err == nil {
				req.ContentLength = contentLength
			}
		}
	}

	return req
}

func (req *HttpRequest) Bytes() []byte {
	var buf bytes.Buffer

	version := "HTTP/1.1"
	if req.Version == 0 {
		version = "HTTP/1.0"
	}

	fmt.Fprintf(
		&buf,
		"%s %s %s\r\n",
		req.Method,
		req.Url,
		version,
	)

	if req.Host != "" {
		fmt.Fprintf(&buf, "Host: %s\r\n", req.Host)
	}

	if req.UserAgent != "" {
		fmt.Fprintf(&buf, "User-Agent: %s\r\n", req.UserAgent)
	}

	if req.ContentType != "" {
		fmt.Fprintf(&buf, "Content-Type: %s\r\n", req.ContentType)
	}

	if req.Body != "" {
		fmt.Fprintf(&buf, "Content-Length: %d\r\n", len(req.Body))
	}

	buf.WriteString("\r\n")

	if req.Body != "" {
		buf.WriteString(req.Body)
	}

	return buf.Bytes()
}
