package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
    "id" string, 
    "name" string,
    "price" float,
    "status" string,
  )`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES(
    "c7bc89c9-3866-4427-9e05-6e92ae9d3ed3", 
    "Test Product",
    0,
    "disabled"
  )`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}
