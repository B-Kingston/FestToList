package main

import (
	"time"
)

// "io"
// "html/template"
// "log"
// "net/http"

func main() {

	// h1 := func(w http.ResponseWriter, r *http.Request) {
	// 	template := template.Must(template.ParseFiles("index.html"))
	// 	template.Execute(w, nil)
	// }
	// http.HandleFunc("/", h1)
	// log.Fatal(http.ListenAndServe(":8000", nil))
	db := initDbConn()
	time_now := time.Now()
	insert(time_now.String(), "fest2list", "submitted_datetime", db)
	// insert(db)
	//setupRoutes()
}
