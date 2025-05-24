package client

import (
	"ZyraCLI/config"
	"bufio"
	"fmt"
	"net"
	"strings"
)

func SendCommand(command string) (string, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return "", fmt.Errorf("config load failed: %w", err)
	}

	conn, err := net.Dial("tcp", cfg.Host)
	if err != nil {
		return "", fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	conn.Write([]byte(fmt.Sprintf("AUTH %s\n", cfg.Token)))
	if line, _ := reader.ReadString('\n'); strings.HasPrefix(line, "ERR") {
		return "", fmt.Errorf("auth failed")
	}

	conn.Write([]byte(command + "\n"))
	response := ""
	for {
		line, err := reader.ReadString('\n')
		if err != nil || strings.TrimSpace(line) == "" {
			break
		}
		response += line
	}
	return response, nil
}