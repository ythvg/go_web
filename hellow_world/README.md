## Introduction

Go is a battery included programming language and has a webserver already build in. The net/http pacakge from the standard library contains all functionalities about the HTTP protocol. This includes(amound many other things) an HTTP client and an HTTP server. In this example you will figure out how simple it is, to create a webserver that you can view in your browser.

## Registering a Request Handler
First, create a Handler which receives all incomming HTTP connections from browsers, HTTP clients or API requests. A handler in Go is a function with signature:
```
func (w http.ResponseWriter, r *http.Request)
```
The function receives two parameters:
An http.ResponseWritter which is where you write your text/html response to.
An http.Request which contains all information about this HTTP request including things like the URL or header fields.
Registering a request handler to the default HTTP Server is as simple as this:
```
http.HandleFunc("/", func (w http.ResponseWritter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
})
```

## Listen for HTTP Connections
The request handler alone can not accept any HTTP connections from the outsides. An HTTP server has to listen on a port to pass connections on to the request handler. Because port 80 is in most cases the default port for HTTP traffic, this server will also listen on it.
The following code will start Go's default HTTP server and listen for connection on port 80. You can navigate your browser to http://localhost/ and see your server handing your request.
```
http.ListenAndServe(":80", nil)
```