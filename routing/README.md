## Introduction
Go's net/http package provides a lot of functionalities for the HTTP protocol. One thing it doesn't do very well is complex request routing like segmenting a request url into single parammeters. Fortunately there is a very popular package for this, which is well know for the good code quality in the Go community. In this example you will see how to use the gorilla/mux package to create routes with named parameters, GET/POST handlers and domain restrictions.

## Installing the gorilla/mux pacakge
gorilla/mux is a package which adapts go Go's default HTTP router. It comes with a lot of features to increase the productivity when writing web applications. It is also compliant to Go's default request handler signature func func (w http.ResponseWriter, r *http.Request), so the package can be mixed and matched with other HTTP libraries like middleware or exising applications. Use the go get command to install the package from GitHub like so:
```
go get -u github.com/gorilla/mux
```

## Creating a new Router
First create a nwe request router. The router is the main router for your web application and will later be passed as parameter to the server. It will receive all HTTP connections and pass it on to the request handlers you will register on it. You can create a new router like so:
```
r := mux.NewRouter()
```

## Registering a Request Handler
Once you have a new router you can register reqeuest handlers like usual. The Only difference is, that instead of calling http.HandleFunc(..l), you call Handle Fucn on your router like this: r.HandleFunc(...).

## URL Parameters
The biggest strength of gorilla/mux Router is the ability to extract segments from the request URL. As an example, this is a URL in your application:
```
/book/go-programming-blueprint/page/10
```
this URL has two dynamic segments:
Book title slug (go-programming-blueprint)
Page (10)
To have a request handler match the URL mentioned above you replace the dynamic segments of with placeholders in your URL pattern like so:
```
r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
    // get the book
    // navigate to the page
})
```
The last thing is to get the data from these segments. The pacakge comes with the function mux.Vars(r) which takes the http.Request as parameter and return a map of the segments.
```
func (w http.ResponseWriter, r *http.Router) {
    vars := mux.Vars(r)
    vars["title"] // the book title slug
    vars["page"] // the page
}
```

## Setting the HTTP server's router
Ever wondered what the nil in http.ListenAndServe(":80", nil) ment? It is the parameter for the main router of the HTTP server. By default it's nil, which means to use the default router of the net/http pacakge. To make use of your own router, replace the nil with the variable of your router r.
```
http.ListenAndServe(":80", r)
```

## Feature of the gorilla/mux Router
## Methods
Restrict the request handler to specific HTTP methods.
```
r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
```

## Hostname & Subdomains
Restrict the request handler to specific hostnames or subdomains.
```
r.HandleFunc("/book/{title}", BookHandler).Host("www.mybookstore.com")
```

## Sechemes
Restrict the request handler to http/https.
```
r.HandleFunc("/secure", SecureHandler).Schemes("https")
r.HandleFunc("/insecure", InsecureHandler).Schemes("http")
```

## Path Prefixes & Subrouters
Restrict the reqeust handler to specific path prefixex.
```
bookrouter := r.PathPrefix("/books").Subrouter()
bookrouter.HandleFunc("/", AllBooks)
bookrouter.HandleFunc("/{title}", GetBook)
```