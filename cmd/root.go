package cmd

import (
	"database/sql"
	"os"

	"github.com/emiliosheinz/exagonal-architecture/adapters/db"
	"github.com/emiliosheinz/exagonal-architecture/application"
	"github.com/spf13/cobra"
)

var database, _ = sql.Open("sqlite3", "sqlite.db")
var productDb = db.NewProductDb(database)
var productService = application.NewProductService(productDb)

var rootCmd = &cobra.Command{
	Use: "exagonal-architecture",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
