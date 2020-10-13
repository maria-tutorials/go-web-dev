package main

import (
	"fmt"
	"log"
	"net/http"

	"./controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()

	r.GET("/", index)

	r.POST("/user", uc.CreateUser)
	r.GET("/user/:id", uc.GetUser)
	r.DELETE("/user/:id", uc.DeleteUser)

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
