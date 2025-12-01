package main

import (
	"fmt"
	"os"
	"pentest-kit/scanner"
	"pentest-kit/utils"
)

func main() {
	var serviceDetection, synScan, aggressive bool
	var host, portRange string

	switch {
	case len(os.Args) == 4 && (os.Args[1] == "-sV" || os.Args[1] == "-sS" || os.Args[1] == "-A"):
		switch os.Args[1] {
		case "-sV":
			serviceDetection = true
		case "-sS":
			synScan = true
		case "-A":
			aggressive = true
		}
		host = os.Args[2]
		portRange = os.Args[3]
	case len(os.Args) == 3:
		host = os.Args[1]
		portRange = os.Args[2]
	default:
		fmt.Println("Usage: go run main.go [-sV|-sS|-A] <host> <port-range>")
		fmt.Println("Example: go run main.go -A 192.168.1.1 80-443")
		os.Exit(1)
	}

	ports := utils.ParsePortRange(portRange)
	if aggressive {
		scanner.AggressiveScan(host, ports)
	} else if synScan {
		scanner.SynScan(host, ports)
	} else {
		scanner.ScanPorts(host, ports, serviceDetection)
	}
}


