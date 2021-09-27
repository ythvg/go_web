package main

import (
	"fmt"
	"log"
	"net/http"
)

// This example will show how to create basic logging middleware in Go.
// A middleware simply takes a http.HandleFunc as one of its parameters, wraps it and returns a new http.HandleFunc for the server to call.

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func main() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))
	http.ListenAndServe(":80", nil)
}
