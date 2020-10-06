package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	oreo, err := req.Cookie("oreo")
	if err == http.ErrNoCookie {
		oreo = &http.Cookie{
			Name:  "oreo",
			Value: "0",
		}
	}

	times, err := strconv.Atoi(oreo.Value)
	if err != nil {
		log.Println("ups", err)
	}

	times++ //always increase
	oreo.Value = strconv.Itoa(times)

	http.SetCookie(w, oreo)

	io.WriteString(w, "Been here: "+oreo.Value)
}
