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


# 3: Handling HTTP response

In the same way we get http response 
For now lets write our own response and send it as bytes

# 4: Concurrency

Nginx can handle thousands of concurrent request due to its event drive architecture.
But for now implementing a event loop will be complex so lets go with handling concurrent request with goroutines
Here the working will be similar
Master Process -> handles creating and managing workers
Worker Process -> Goroutines handling http request


# 5: Serving static files 

Nginx also can server static files
If the http request is asking for a static file and if it exist in the system then return it


# 6: Reverse Proxy

Nginx does reverse proxy by checking against the config and if the url matches
it send the http req where it can add some headers and then sent the req to url backend mentioned in config
the response is received and convert to a standard response and returned to the client 
here using http lib for easy handling 


# 7: Loadbalancer 

Nginx also does loadbalancing. It distributes load to different servers.
which is configured in the config file.
Here i implemented a simple loadbalancer using round robin