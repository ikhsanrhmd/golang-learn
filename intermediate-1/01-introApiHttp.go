package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type articel struct {
	ID    int
	Title string
	Nama  string
}

var data = []articel{
	articel{1, "mr.", "rahmad"},
	articel{2, "mr.", "ilham"},
}

func articels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)

}
func main() {
	http.HandleFunc("/artikel", articels)
	fmt.Println("starting web at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
