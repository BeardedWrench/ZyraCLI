package cmd

import (
	"ZyraCLI/internal"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var createTableCmd = &cobra.Command{
	Use:   "create-table",
	Short: "Create a new table in ZyraDB",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Enter table name: ")
		tableName, _ := r.ReadString('\n')
		tableName = strings.TrimSpace(tableName)

		fmt.Println("Paste JSON schema (fields only, no name). End with a line containing only 'EOF':")
		var lines []string
		for {
			line, _ := r.ReadString('\n')
			if strings.TrimSpace(line) == "EOF" {
				break
			}
			lines = append(lines, strings.TrimSpace(line))
		}

		raw := strings.Join(lines, "")
		var buf bytes.Buffer
		err := json.Compact(&buf, []byte(raw))
		if err != nil {
			return fmt.Errorf("❌ Invalid JSON input: %w", err)
		}

		createCmd := fmt.Sprintf("CREATE_TABLE %s {\"fields\":%s}", tableName, buf.String())

		conn, err := internal.ConnectAndAuth()
		if err != nil {
			return err
		}
		defer conn.Close()

		fmt.Println("🔧 Full CREATE_TABLE command:")
		fmt.Println(createCmd)
		return internal.SendCommandAndPrint(conn, createCmd)
	},
}

func init() {
	rootCmd.AddCommand(createTableCmd)
}