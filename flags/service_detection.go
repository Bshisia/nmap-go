package flags

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func DetectService(conn net.Conn, port int) string {
	services := map[int]string{
		22: "ssh", 23: "telnet", 25: "smtp", 53: "dns", 80: "http",
		110: "pop3", 143: "imap", 443: "https", 993: "imaps", 995: "pop3s",
	}
	
	if service, exists := services[port]; exists {
		if port == 80 || port == 443 {
			fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
			conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			scanner := bufio.NewScanner(conn)
			if scanner.Scan() {
				response := scanner.Text()
				if strings.Contains(response, "HTTP") {
					return service
				}
			}
		}
		return service
	}
	return "unknown"
}