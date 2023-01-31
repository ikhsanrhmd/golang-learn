package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you requested: %s, with token %s", r.URL.Path, r.URL.Query().Get("token"))
}

func main() {
	http.HandleFunc("/", rootHandler)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("WEB STARTED")
	http.ListenAndServe(":80", nil)
}
