package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string

	if req.Method == http.MethodPost {
		s = handleFile(w, req)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`)
}

func handleFile(w http.ResponseWriter, req *http.Request) string {
	// open
	f, h, err := req.FormFile("q")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}
	defer f.Close()

	// for your information
	fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

	// read
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}

	storeFile(w, h.Filename, bs)

	return string(bs)
}

func storeFile(w http.ResponseWriter, filename string, bs []byte) {
	// store on server
	dst, err := os.Create(filepath.Join("./user/", filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = dst.Write(bs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
