package main

import (
	"database/sql"
	"fmt"

	"github.com/emiliosheinz/exagonal-architecture/adapters/db"
	"github.com/emiliosheinz/exagonal-architecture/application"
)

func main() {
	database, _ := sql.Open("sqlite3", "sqlite.db")
	database.Exec(`CREATE TABLE IF NOT EXISTS products (id string, name string, price float, status string)`)

	productDbAdapter := db.NewProductDb(database)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Test Product", 25.0)
	fmt.Println("Product created: ")
	fmt.Println(product)

	productService.Enable(product)
	fmt.Println("Product enabled: ")
	fmt.Println(product)
}
