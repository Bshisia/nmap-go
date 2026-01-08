package main

import (
	"fmt"
	"os"
	"pentest-kit/scanner"
	"pentest-kit/utils"
)

func main() {
	var serviceDetection, synScan, aggressive, udpScan, finScan, xmasScan, nullScan bool
	var host, portRange string

	switch {
	case len(os.Args) == 4 && (os.Args[1] == "-sV" || os.Args[1] == "-sS" || os.Args[1] == "-A" || os.Args[1] == "-sU" || os.Args[1] == "-sF" || os.Args[1] == "-sX" || os.Args[1] == "-sN"):
		switch os.Args[1] {
		case "-sV":
			serviceDetection = true
		case "-sS":
			synScan = true
		case "-A":
			aggressive = true
		case "-sU":
			udpScan = true
		case "-sF":
			finScan = true
		case "-sX":
			xmasScan = true
		case "-sN":
			nullScan = true
		}
		host = os.Args[2]
		portRange = os.Args[3]
	case len(os.Args) == 3:
		host = os.Args[1]
		portRange = os.Args[2]
	default:
		fmt.Println("Usage: go run main.go [-sV|-sS|-A|-sU|-sF|-sX|-sN] <host> <port-range>")
		fmt.Println("Example: go run main.go -A 192.168.1.1 80-443")
		os.Exit(1)
	}

	ports := utils.ParsePortRange(portRange)
	if aggressive {
		scanner.AggressiveScan(host, ports)
	} else if synScan {
		scanner.SynScan(host, ports)
	} else if udpScan {
		scanner.UdpScan(host, ports)
	} else if finScan {
		scanner.FinScan(host, ports)
	} else if xmasScan {
		scanner.XmasScan(host, ports)
	} else if nullScan {
		scanner.NullScan(host, ports)
	} else {
		scanner.ScanPorts(host, ports, serviceDetection)
	}
}


