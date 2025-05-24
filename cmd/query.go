package cmd

import (
	"ZyraCLI/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query [table]",
	Short: "Query all data in a table",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := internal.ConnectAndAuth()
		if err != nil {
			return err
		}
		defer conn.Close()
		return internal.SendCommandAndPrint(conn, fmt.Sprintf("QUERY %s", args[0]))
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}