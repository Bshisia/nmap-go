package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

