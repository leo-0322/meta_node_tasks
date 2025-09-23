package main

import "github.com/jmoiron/sqlx"

var db *sqlx.DB

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	var book []Book

	query := "selct * from books where price > ?"
	err := db.Select(&book, query, 50)
	if err != nil {
		panic(err)
	}
}
