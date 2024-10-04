package cmd

import (
	"fmt"

	"github.com/emiliosheinz/exagonal-architecture/adapters/web/server"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use: "http",
	Run: func(cmd *cobra.Command, args []string) {
		webserver := server.NewWebserver()
		webserver.Service = productService
		fmt.Println("Webserver has been started")
		webserver.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
