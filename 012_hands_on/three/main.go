//Create a data structure to pass to a template which
//contains information about restaurant's menu including Breakfast, Lunch, and Dinner items
//for an unlimited number of restaurants
// also add a function to make something uppercase
package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"upper": strings.ToUpper,
}

type meal struct {
	Name, Description string
	Price             int
}

type menu []meal

type restaurant struct {
	Name string
	Menu menu
}

type restaurants []restaurant

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("tpl.gohtml"))
}

func main() {
	r := restaurants{
		restaurant{
			"Bon appetit",
			menu{
				meal{
					"CheeseHamburguer",
					"Burguer that has cheese",
					10,
				},
				meal{
					"Salad",
					"Green things, healthy",
					11,
				},
			},
		},
		restaurant{
			"Pasta Pasta Pasta",
			menu{
				meal{
					"Carbonara",
					"Spaghetti",
					12,
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", r)
	if err != nil {
		log.Fatalln(err)
	}
}
