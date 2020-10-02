package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	landing := Page{
		Title:   "Escaped",
		Heading: "Danger is escaped with html/template",
		Input:   `<script>alert("Yow!");</script>`,
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("ups...", err)
	}
	defer nf.Close()

	err = tpl.ExecuteTemplate(nf, "index.gohtml", landing)
	if err != nil {
		log.Fatalln(err)
	}
}
