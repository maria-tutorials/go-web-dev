package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")
	check(err)
	defer db.Close()

	//	if err := db.Ping(); err != nil {
	//		log.Fatal(err)
	//	}
	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "Yey did it.")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
