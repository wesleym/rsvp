package main

import (
	"log"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", Root)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
