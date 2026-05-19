package proxy

import (
	"bytes"
	"io"
	nethttp "net/http"
	"nigx/internals/http"
	"strings"
)

func ProxyRequest(req *http.HttpRequest, proxyUrl string, params string) []byte {
	url := proxyUrl + params
	client := &nethttp.Client{}
	proxyReq, _ := nethttp.NewRequest(
		string(req.Method),
		url,
		bytes.NewReader([]byte(req.Body)),
	)

	resp, err := client.Do(proxyReq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	headers := map[string]string{}

	for k, v := range resp.Header {
		headers[k] = strings.Join(v, ", ")
	}
	httpresp := http.NewHttpResponse("1.1", resp.StatusCode, nethttp.StatusText(resp.StatusCode), headers, body)
	return httpresp.Bytes()
}

func IsProxyRequest(route, url string) (string, bool) {
	if strings.HasPrefix(url, route) {
		rest := strings.TrimPrefix(url, route)
		return rest, true
	}
	return "", false
}
