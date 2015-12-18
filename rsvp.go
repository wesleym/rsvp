package main

import (
	"html/template"
	"log"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Root)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
