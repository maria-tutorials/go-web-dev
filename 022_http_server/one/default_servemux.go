package main

import (
	"io"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "We are at /")
}

func meHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "My name's Maria :)")
}

func dogHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "We are at /dog")
}

func main() {

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/me", meHandler)

	http.HandleFunc("/dog", dogHandler)

	http.ListenAndServe(":8080", nil)
}
