package main

import (
	"io"
	"net/http"
)

type pet int

func (m pet) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "puppy puppy puppy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	var d pet
	http.ListenAndServe(":8080", d)
}
