package main

import (
	"fmt"
	"os"
	"pentest-kit/scanner"
	"pentest-kit/utils"
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

	ports := utils.ParsePortRange(portRange)
	scanner.ScanPorts(host, ports, serviceDetection)
}


