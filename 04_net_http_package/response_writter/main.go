package main

import (
	"fmt"
	"net/http"
)

type icecream int

func (m icecream) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Maria-Key", "this is from maria")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in here</h1>")
}

func main() {
	var d icecream
	http.ListenAndServe(":8080", d)
}
