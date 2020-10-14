package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Book order and naming matches the table
type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	db, err := sql.Open("postgres", "postgres://maria:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	rows, err := db.Query("SELECT * FROM books;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, bk := range bks {
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}

// psql
// create database bookstore;
// create user maria with password 'password';
// grant all privileges on database bookstore to maria
