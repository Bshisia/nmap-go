package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [-sV] <host> <port-range>")
		fmt.Println("Example: go run main.go -sV 192.168.1.1 80-443")
		os.Exit(1)
	}

	var serviceDetection bool
	var host, portRange string

	if os.Args[1] == "-sV" {
		serviceDetection = true
		host = os.Args[2]
		portRange = os.Args[3]
	} else {
		host = os.Args[1]
		portRange = os.Args[2]
	}

	ports := parsePortRange(portRange)
	scanPorts(host, ports, serviceDetection)
}

func parsePortRange(portRange string) []int {
	var ports []int
	
	if strings.Contains(portRange, "-") {
		parts := strings.Split(portRange, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		
		for i := start; i <= end; i++ {
			ports = append(ports, i)
		}
	} else {
		port, _ := strconv.Atoi(portRange)
		ports = append(ports, port)
	}
	
	return ports
}

func scanPorts(host string, ports []int, serviceDetection bool) {
	fmt.Printf("Scanning %s...\n", host)
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		
		if err == nil {
			if serviceDetection {
				service := detectService(conn, port)
				fmt.Printf("Port %d: OPEN %s\n", port, service)
			} else {
				fmt.Printf("Port %d: OPEN\n", port)
			}
			conn.Close()
		}
	}
}

func detectService(conn net.Conn, port int) string {
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