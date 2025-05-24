package cmd

import (
	"ZyraCLI/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var dropTableCmd = &cobra.Command{
	Use:   "drop-table [table]",
	Short: "Drop a table",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := internal.ConnectAndAuth()
		if err != nil {
			return err
		}
		defer conn.Close()
		return internal.SendCommandAndPrint(conn, fmt.Sprintf("DROP_TABLE %s", args[0]))
	},
}

func init() {
	rootCmd.AddCommand(dropTableCmd)
}