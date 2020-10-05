package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	fs := http.FileServer(http.Dir("."))

	http.Handle("/public/", fs)

	http.Handle("/resources/", http.StripPrefix("/resources", fs))

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/sample/", sample)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func sample(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "sample.gohtml", "Maria")
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
