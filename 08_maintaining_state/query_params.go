package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")

	if v == "" {
		io.WriteString(w, "no query passed for q")
	} else {
		fmt.Fprintln(w, "my search: "+v)
	}

}
