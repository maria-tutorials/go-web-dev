package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	plants := []string{"Calathea", "Begonia", "Fitonia", "Monstera"}

	err := tpl.Execute(os.Stdout, plants)
	if err != nil {
		log.Fatalln(err)
	}

}
