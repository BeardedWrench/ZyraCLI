package cmd

import (
	"ZyraCLI/internal"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var insertCmd = &cobra.Command{
	Use:   "insert [table]",
	Short: "Insert a record into a table",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Paste JSON object:")
		obj, _ := reader.ReadString('\n')
		obj = strings.TrimSpace(obj)

		conn, err := internal.ConnectAndAuth()
		if err != nil {
			return err
		}
		defer conn.Close()
		return internal.SendCommandAndPrint(conn, fmt.Sprintf("INSERTJSON %s %s", args[0], obj))
	},
}


func init() {
	rootCmd.AddCommand(insertCmd)
}