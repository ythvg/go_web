package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	br := r.PathPrefix("/books").Subrouter()
	br.HandleFunc("/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "you've reqeusted the book: %s on page %s\n", title, page)
	})
	http.ListenAndServe(":80", r)
}
