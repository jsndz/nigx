# 1 : TCP connections

Nginx is built on top of tcp with its custom implementation for HTTP
It is highly optimised because of this
So start with a tcp server that accepts message and returns response.

# 2: Handling HTTP requests

So basically when we do a http req what tcp server recieves is this:

```curl
POST / HTTP/1.1
Host: localhost:8080
User-Agent: curl/7.81.0
Accept: */*
Content-Type: application/json
Content-Length: 19

{"message":"hello"}
```

this in bytes
Here we can parse this to get a http request