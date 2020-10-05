package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/puppy", puppyHandler)
	http.HandleFunc("/pictures/puppy.jpg", puppyPictureHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "index function ran")
}

func puppyHandler(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func puppyPictureHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "pictures/puppy.jpg")
}
