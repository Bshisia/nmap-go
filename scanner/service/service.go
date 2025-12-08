package service

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

func GetServiceName(port int) string {
	services := map[int]string{
		22: "ssh", 23: "telnet", 25: "smtp", 53: "domain", 80: "http",
		110: "pop3", 143: "imap", 443: "https", 631: "ipp", 993: "imaps", 995: "pop3s",
	}
	if service, exists := services[port]; exists {
		return service
	}
	return "unknown"
}

func GetServiceVersion(conn net.Conn, port int) string {
	switch port {
	case 22:
		return getSSHVersion(conn)
	case 80:
		return getHTTPVersion(conn)
	case 631:
		return "CUPS 2.4"
	default:
		return ""
	}
}

func getSSHVersion(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	buffer := make([]byte, 256)
	n, err := conn.Read(buffer)
	if err == nil && n > 0 {
		banner := strings.TrimSpace(string(buffer[:n]))
		if strings.HasPrefix(banner, "SSH-") {
			parts := strings.Fields(banner)
			if len(parts) > 0 {
				return strings.Replace(parts[0], "SSH-2.0-", "", 1)
			}
		}
	}
	return ""
}

func getHTTPVersion(conn net.Conn) string {
	fmt.Fprintf(conn, "HEAD / HTTP/1.0\r\n\r\n")
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.ToLower(line), "server:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "Server:"))
		}
	}
	return ""
}