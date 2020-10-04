package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "tpl.gohtml", "Maria")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func meHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "My name's Maria :)")
}

func dogHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "We are at /dog")
}

func main() {

	http.HandleFunc("/", indexHandler)

	http.Handle("/me", http.HandlerFunc(meHandler))

	http.HandleFunc("/dog", dogHandler)

	http.ListenAndServe(":8080", nil)
}
