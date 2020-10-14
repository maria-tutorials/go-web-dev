package main

import (
	"net/http"

	"./controllers"
)

// Book order and naming matches the table
type Book struct {
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/books", controllers.Index)
	http.HandleFunc("/books/show", controllers.Show)
	http.HandleFunc("/books/create", controllers.Create)
	http.HandleFunc("/books/create/process", controllers.CreateProcess)
	http.HandleFunc("/books/update", controllers.Update)
	http.HandleFunc("/books/update/process", controllers.UpdateProcess)
	http.HandleFunc("/books/delete/process", controllers.DeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
