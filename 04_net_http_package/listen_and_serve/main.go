package main

import (
	"fmt"
	"net/http"
)

type icecream int

func (m icecream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "We still don't have an HandleFunc so this will do")
}

func main() {
	var d icecream
	http.ListenAndServe(":8080", d)
}
