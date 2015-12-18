package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"fmt"

	_ "github.com/lib/pq"
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
	p := os.Getenv("POSTGRES_ENV_POSTGRES_PASSWORD")
	s := fmt.Sprintf("user=postgres dbname=postgres password=%s sslmode=verify-full", p)
	_, err := sql.Open("postgres", s)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", Root)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
