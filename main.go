package main

import (
	"fmt"
	"os"
	"pentest-kit/scanner"
	"pentest-kit/utils"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [-sV|-sS] <host> <port-range>")
		fmt.Println("Example: go run main.go -sS 192.168.1.1 80-443")
		os.Exit(1)
	}

	var serviceDetection, synScan bool
	var host, portRange string

	switch os.Args[1] {
	case "-sV":
		serviceDetection = true
		host = os.Args[2]
		portRange = os.Args[3]
	case "-sS":
		synScan = true
		host = os.Args[2]
		portRange = os.Args[3]
	default:
		host = os.Args[1]
		portRange = os.Args[2]
	}

	ports := utils.ParsePortRange(portRange)
	if synScan {
		scanner.SynScan(host, ports)
	} else {
		scanner.ScanPorts(host, ports, serviceDetection)
	}
}


