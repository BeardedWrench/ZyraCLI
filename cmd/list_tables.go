package cmd

import (
	"ZyraCLI/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var listTablesCmd = &cobra.Command{
	Use:   "list-tables",
	Short: "List tables from ZyraDB over TCP",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Using ConnectAndAuth from internal/client.go")
		
		conn, err := internal.ConnectAndAuth()
		if err != nil {
			return err
		}
		defer conn.Close()

		return internal.SendCommandAndPrint(conn, "LIST_TABLES")
	},
}

func init() {
	rootCmd.AddCommand(listTablesCmd)
}