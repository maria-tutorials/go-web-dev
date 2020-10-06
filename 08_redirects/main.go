package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/seeOther", seeOtherHandler)
	http.HandleFunc("/tempRedirect", tempRedirectHandler)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at /: ", req.Method, "\n\n")
}

func seeOtherHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at /seeOther:", req.Method)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func tempRedirectHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at /tempRedirect:", req.Method)
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
