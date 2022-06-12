package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> Hello Welcome!! </h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact me</h1>")
}

func urlHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		homeHandler(w, r)
	} else if r.URL.Path == "/contact" {
		contactHandler(w, r)
	} else {
		http.Error(w, "404 Not Found", http.StatusNotFound) //use the default http.Error of go
	}
}

func main() {
	r := chi.NewRouter()
	fmt.Println("Listening on server :7000")
	http.ListenAndServe(":7000", r)
}
