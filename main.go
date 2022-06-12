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

//either this code can be implimented as else handler function
// func notFounError(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<h1>Error</h1>")
// }

func urlHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		homeHandler(w, r)
	} else if r.URL.Path == "/contact" {
		contactHandler(w, r)
	} else {
		// 	w.WriteHeader(http.StatusNotFound) //manually set the header and message
		// 	fmt.Fprintf(w, "Page not found")
		http.Error(w, "404 Not Found", http.StatusNotFound) //use the default http.Error of go
	}
}

func main() {
	// http.HandleFunc("/", urlHandler) // this has been removed ad the path handler can be directly cast to listen and serrve
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Listening on server :7000")
	http.ListenAndServe(":7000", http.HandlerFunc(urlHandler))
}
