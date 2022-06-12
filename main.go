package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> Hello Welcome!! </h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact me</h1>")
}

func notFounError(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Error</h1>")
}

func urlHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		homeHandler(w, r)
	} else if r.URL.Path == "/contact" {
		contactHandler(w, r)
	} else {
		notFounError(w, r)
	}
}

func main() {
	http.HandleFunc("/", urlHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Listening on server :3000")
	http.ListenAndServe(":3000", nil)
}
