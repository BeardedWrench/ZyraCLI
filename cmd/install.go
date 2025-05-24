package cmd

import (
	"ZyraCLI/config"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Setup CLI config for ZyraDB",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Enter ZyraDB host (e.g. 127.0.0.1): ")
		host, _ := r.ReadString('\n')
		fmt.Print("Enter ZyraDB port (e.g. 7070): ")
		port, _ := r.ReadString('\n')
		fmt.Print("Enter API token (optional): ")
		token, _ := r.ReadString('\n')

		cfg := config.Config{
			Host:  strings.TrimSpace(host),
			Port:  strings.TrimSpace(port),
			Token: strings.TrimSpace(token),
		}
		return config.SaveConfig(cfg)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}