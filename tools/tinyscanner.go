package tools

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func TinyScanner(target string, ports string, output string) {
	portList := parsePorts(ports)
	results := []string{}
	
	results = append(results, fmt.Sprintf("Scanning target: %s", target))
	results = append(results, fmt.Sprintf("Ports: %s", ports))
	results = append(results, "")
	
	for _, port := range portList {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		
		if err == nil {
			result := fmt.Sprintf("Port %d: OPEN", port)
			results = append(results, result)
			fmt.Println(result)
			conn.Close()
		} else {
			result := fmt.Sprintf("Port %d: CLOSED", port)
			results = append(results, result)
			fmt.Println(result)
		}
	}
	
	if output != "" {
		writeToFile(output, results)
	}
}

func parsePorts(ports string) []int {
	portList := []int{}
	parts := strings.Split(ports, ",")
	
	for _, part := range parts {
		var port int
		fmt.Sscanf(part, "%d", &port)
		if port > 0 && port <= 65535 {
			portList = append(portList, port)
		}
	}
	
	return portList
}

func writeToFile(filename string, lines []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()
	
	for _, line := range lines {
		file.WriteString(line + "\n")
	}
	fmt.Printf("\nResults saved to %s\n", filename)
}
