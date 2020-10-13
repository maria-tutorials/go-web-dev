package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.GET("/", index)
	r.GET("/user/:id", getUser)

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Index</title>
			</head>
			<body>
			<a href="/user/123456789">CHECK OUT: http://localhost:8080/user/123456789</a>
			</body>
			</html>
		`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		ID:     p.ByName("id"),
		Name:   "FAKE USER",
		Gender: "fake",
		Age:    100,
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ups something went wrong")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}
