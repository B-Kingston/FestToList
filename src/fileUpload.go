package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File")
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	var fileBytes []byte
	fileBytes, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, " Successfully Uploaded File\n")
	println(fileBytes)
}

func displayHomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	test := "test"
	err = t.Execute(w, test)
	if err != nil {
		fmt.Println(err)
	}
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/home", displayHomePage)
	http.ListenAndServe(":8080", nil)
}
