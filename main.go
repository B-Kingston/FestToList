package main

import (
	// "fmt"
	// "io"
	"html/template"
	"log"
	"net/http"
)

type Search struct {
	QueryId       int
	QueryUpload   string
	QueryDateTime string
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("index.html"))
		template.Execute(w, nil)
	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
