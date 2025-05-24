package cmd

import (
	"ZyraCLI/internal"
	"bufio"
	"fmt"

	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Manage API tokens remotely",
}

var tokenAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new API token",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := internal.ConnectAndAuth()
		if err != nil {
			return err
		}
		defer conn.Close()

		fmt.Fprintln(conn, "ADD_TOKEN")
		resp, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(resp)
		return nil
	},
}

var tokenRemoveCmd = &cobra.Command{
	Use:   "remove [token]",
	Short: "Remove an API token",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := internal.ConnectAndAuth()
		if err != nil {
			return err
		}
		defer conn.Close()

		fmt.Fprintf(conn, "REMOVE_TOKEN %s\n", args[0])
		resp, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(resp)
		return nil
	},
}


var tokenListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all API tokens",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := internal.ConnectAndAuth()
		if err != nil {
			return err
		}
		defer conn.Close()

		fmt.Fprintln(conn, "LIST_TOKENS")
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		return scanner.Err()
	},
}

func init() {
	tokenCmd.AddCommand(tokenAddCmd)
	tokenCmd.AddCommand(tokenRemoveCmd)
	tokenCmd.AddCommand(tokenListCmd)
	rootCmd.AddCommand(tokenCmd)
}