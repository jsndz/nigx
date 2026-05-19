package http

import (
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

	// Parse request line
	// Example: POST / HTTP/1.1
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
