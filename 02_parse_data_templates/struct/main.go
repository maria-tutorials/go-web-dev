package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type plant struct {
	Name string
	Type string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p := plant{
		"Begonia",
		"maculata",
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}

}
