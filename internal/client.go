package internal

import (
	"ZyraCLI/config"
	"bufio"
	"fmt"
	"net"
	"strings"
)
func SendCommandAndPrint(conn net.Conn, cmd string) error {
	fmt.Println("➡️ Sending command:", cmd)
	_, err := fmt.Fprintln(conn, cmd)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("❌ Failed to read response: %w", err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		fmt.Println("🟢 Response:", line)

		if line == "OK" || strings.HasPrefix(line, "ERR") {
			break
		}
	}
	return nil
}

func ConnectAndAuth() (net.Conn, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	fmt.Printf("DEBUG CONFIG: host=%q port=%q token=%q\n", cfg.Host, cfg.Port, cfg.Token)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		return nil, err
	}
	fmt.Println("✅ Auth successful, connection ready")
	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	fmt.Println("Server says:", strings.TrimSpace(line))
	fmt.Fprintln(conn, "AUTH "+cfg.Token)
	resp, _ := r.ReadString('\n')
	if strings.HasPrefix(resp, "ERR") {
		return nil, fmt.Errorf("auth failed: %s", resp)
	}
	return conn, nil
}

