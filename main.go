package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <host> <port-range>")
		fmt.Println("Example: go run main.go 192.168.1.1 80-443")
		os.Exit(1)
	}

	host := os.Args[1]
	portRange := os.Args[2]

	ports := parsePortRange(portRange)
	scanPorts(host, ports)
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

func scanPorts(host string, ports []int) {
	fmt.Printf("Scanning %s...\n", host)
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		
		if err == nil {
			fmt.Printf("Port %d: OPEN\n", port)
			conn.Close()
		}
	}
}