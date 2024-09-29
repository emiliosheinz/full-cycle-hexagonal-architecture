package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/emiliosheinz/exagonal-architecture/adapters/db"
	"github.com/emiliosheinz/exagonal-architecture/application"
	"github.com/stretchr/testify/require"
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
    "status" string
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

func TestProductDb_Get(t *testing.T) {
	setup()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("c7bc89c9-3866-4427-9e05-6e92ae9d3ed3")
	require.Nil(t, err)
	require.Equal(t, "Test Product", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestSaveProductDb_Save(t *testing.T) {
	setup()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Test Product"
	product.Price = 100.0

	result, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetName(), result.GetName())
	require.Equal(t, product.GetPrice(), result.GetPrice())
	require.Equal(t, product.GetStatus(), result.GetStatus())

	product.Status = "enabled"

	result, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetName(), result.GetName())
	require.Equal(t, product.GetPrice(), result.GetPrice())
	require.Equal(t, product.GetStatus(), result.GetStatus())
}
