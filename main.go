package main

import (
	"fmt"
	"pentest-kit/scanner"
	"pentest-kit/utils"
)

func main() {
	config := ParseArgs()
	ports := utils.ParsePortRange(config.PortRange)

	if config.Verbose {
		fmt.Printf("Starting scan with verbose output...\n")
	}
	if config.Timing != "" {
		fmt.Printf("Using timing template: %s\n", config.Timing)
	}

	switch {
	case config.Aggressive:
		scanner.AggressiveScan(config.Host, ports)
	case config.SynScan:
		scanner.SynScan(config.Host, ports)
	case config.UdpScan:
		scanner.UdpScan(config.Host, ports)
	case config.FinScan:
		scanner.FinScan(config.Host, ports)
	case config.XmasScan:
		scanner.XmasScan(config.Host, ports)
	case config.NullScan:
		scanner.NullScan(config.Host, ports)
	case config.OSDetection:
		scanner.OSDetection(config.Host, ports)
	default:
		scanner.ScanPorts(config.Host, ports, config.ServiceDetection)
	}
}


